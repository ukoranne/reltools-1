import os

MODEL_BASE_PATH = "/home/ccordes/git/external/src/github.com/openconfig/release/models"

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
        if "fmt" in name and \
           ".go" in name:
            if os.path.isfile(path) and "enum" not in path and "func" not in path:
                yield (dir, name)
        elif not "." in name:
            for d, f  in scan_dir_for_go_files(path):
                yield (d, f)

def build_thrift_from_go():
    # generate thrift code from go code
    for dir, gofilename in scan_dir_for_go_files(MODEL_BASE_PATH):
        print dir, gofilename, dir.split('/')[-1]
        thriftFileName = dir.split('/')[-1] + "d.thrift"
        thriftfd = open(thriftFileName, 'w')

        path = os.path.join(dir, gofilename)
        gofd = open(path, 'r')
        thriftfd.write("namespace go %s" %(dir.strip('/')[-1]+'d'))
        deletingComment = False
        for line in gofd.readline():
            if not deletingComment:
                if "struct" in line or \
                   "}" in line:
                    thriftfd.write(line)
                # lets skip all blank lines
                # skip comments
                elif line == '\n' or \
                   "#" in line or \
                   "package" in line:
                    continue
                elif "/*" in line:
                    deletingComment = True
                else: # found element in struct
                    print "found element line", line
                    elemtype = line.strip(' ')[-1]
                    print elemtype
                    if elemtype in goToThirftTypeMap:
                        thriftfd.write(line.strip(' ')[:-1]+goToThirftTypeMap[elemtype]["native_type"])

            else:
                if "*/" in line:
                    deletingComment = false
                continue
        thriftfd.close()

if __name__ == "__main__":

    build_thrift_from_go()