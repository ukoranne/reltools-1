import os
import json
import pprint

HOME = os.getenv("HOME")
GO_MODEL_BASE_PATH = HOME + "/git/snaproute/generated/src/gomodel/"
JSON_MODEL_REGISTRAION_PATH = HOME + "/git/snaproute/src/models/"

goToThirftTypeMap = {
  'bool':          {"native_type": "bool"},
  'uint8':            {"native_type": "byte", "unsigned": True},
  'uint16':           {"native_type": "i16", "unsigned": True},
  'uint32':           {"native_type": "i32", "unsigned": True},
  'uint64':           {"native_type": "i64", "unsigned": True},
  'string':           {"native_type": "string","unsigned": None },
  'float64':          {"native_type": "double", "unsigned": False},
  'int8':             {"native_type": "byte", "unsigned": False},
  'int16':            {"native_type": "i16", "unsigned": False},
  'int32':            {"native_type": "i32", "unsigned": False},
  'int64':            {"native_type": "i64", "unsigned": False},
}



def scan_dir_for_go_files(dir):
    for name in os.listdir(dir):
        #print "x", dir, name
        path = os.path.join(dir, name)
        if name.endswith('.go'):
            if os.path.isfile(path) and "enum" not in path and "func" not in path:
                yield (dir, name)
        elif not "." in name:
            for d, f  in scan_dir_for_go_files(path):
                yield (d, f)

def scan_dir_for_json_files(dir):
    for name in os.listdir(dir):
        #print "x", dir, name
        path = os.path.join(dir, name)
        if name.endswith('.json'):
            if os.path.isfile(path):
                yield (dir, name)
        elif not "." in name:
            for d, f  in scan_dir_for_go_files(path):
                yield (d, f)

def build_thrift_from_go():
    # generate thrift code from go code
    goStructToListersDict = {}
    deamons = []
    for dir, gofilename in scan_dir_for_json_files(JSON_MODEL_REGISTRAION_PATH):
        path = os.path.join(dir, gofilename)
        print path
        with open(path, 'r') as f:
            data = json.load(f)

            for k, v in data.iteritems():
                if v["Owner"]:
                    goStructToListersDict.setdefault(k, [])
                    goStructToListersDict[k].append(v["Owner"])
                    if v["Owner"] not in deamons:
                        deamons.append(v["Owner"])
                if v["Listeners"]:
                    goStructToListersDict.setdefault(k, [])
                    goStructToListersDict[k] += v["Listeners"]
                    for d  in v["Listeners"]:
                        if d not in deamons:
                            deamons.append(d)
    pprint.pprint(goStructToListersDict)
    print deamons
    for d in deamons:
        thriftFileName = d + ".thrift"
        thriftfd = open(thriftFileName, 'w')
        thriftfd.write("namespace go %s\n" %(d))
        crudStructsList = []

        for dir, gofilename in scan_dir_for_go_files(GO_MODEL_BASE_PATH):
            print dir, gofilename, dir.split('/')[-1]

            path = os.path.join(dir, gofilename)
            gofd = open(path, 'r')
            deletingComment = False
            writingStruct = False
            memberCnt = 1
            for line in gofd.readlines():
                if not deletingComment:
                    if "struct" in line:
                        lineSplit = line.split(" ")
                        structLine = "struct " + lineSplit[1] + "{\n"
                        if lineSplit[1] in goStructToListersDict:
                            print "found line now checking deamon", d, goStructToListersDict[lineSplit[1]]
                            if d in goStructToListersDict[lineSplit[1]]:
                                thriftfd.write(structLine)
                                crudStructsList.append(lineSplit[1])
                                writingStruct = True
                    elif "}" in line and writingStruct:
                        thriftfd.write("}\n")
                        writingStruct = False
                        memberCnt = 1
                    # lets skip all blank lines
                    # skip comments
                    elif line == '\n' or \
                        "//" in line or \
                       "#" in line or \
                       "package" in line:
                        continue
                    elif "/*" in line:
                        deletingComment = True
                    elif writingStruct: # found element in struct
                        #print "found element line", line
                        lineSplit = line.split(' ')
                        print lineSplit
                        elemtype = lineSplit[-3].rstrip('\n') if 'KEY' in lineSplit[-1] else lineSplit[-1].rstrip('\n')
                        #print lineSplit
                        print elemtype, type(elemtype), goToThirftTypeMap.keys()
                        if elemtype in goToThirftTypeMap.keys():
                            print "writing element"
                            thriftfd.write("\t%s : %s %s\n" %(memberCnt,
                                                            goToThirftTypeMap[elemtype]["native_type"],
                                                            lineSplit[0]))
                        memberCnt += 1
                else:
                    if "*/" in line:
                        deletingComment = False

        print crudStructsList
        thriftfd.write("service %sServer {\n" %(d.upper()))
        for s in crudStructsList:
            thriftfd.write("""\tCreate%s(1:%s %s);\n\tUpdate%s(1:%s %s);\n\tDelete%s(1:%s %s);\n\n""" %(s, s, s.lower(), s, s, s.lower(), s, s, s.lower(),) )
        thriftfd.write("}")
        thriftfd.close()

if __name__ == "__main__":

    build_thrift_from_go()