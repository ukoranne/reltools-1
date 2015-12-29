import os
import subprocess
import json
import pprint

OBJECT_MAP_NAME = "objectmap.go"

MODEL_NAME = 'models'
srBase = os.environ.get('SR_CODE_BASE', None)
GO_MODEL_BASE_PATH = srBase + "/generated/src/%s/" % MODEL_NAME
JSON_MODEL_REGISTRAION_PATH = srBase + "/snaproute/src/models/"
#JSON_MODEL_REGISTRAION_PATH = HOME + "/git/reltools/codegentools/gotojson/"
CODE_GENERATION_PATH = srBase + "/reltools/codegentools/gotothrift/"
CLIENTIF_CODE_GENERATION_PATH = srBase + "/generated/src/config/"
OBJMAP_CODE_GENERATION_PATH = srBase + "/generated/src/%s/" % MODEL_NAME
THRIFT_CODE_GENERATION_PATH = srBase + "/generated/src/gorpc/"

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

gDryRun =  False
def executeGoFmtCommand (fd, command, dstPath) :
    out = ''
    if type(command) != list:
        command = [ command]
    for cmd in command:
        if gDryRun :
            print cmd
        else:
            process = subprocess.Popen(cmd.split(), stdout=subprocess.PIPE)
            out,err = process.communicate()
            # create a go format version, at this point the fd is still
            # open so this is a .tmp file, lets strip this for the new
            # file
            #print out, err
            dir = CODE_GENERATION_PATH
            fmt_name_with_dir = dir + fd.name
            #print fmt_name_with_dir
            if not os.path.exists(dir):
              os.makedirs(dir)
            #nfd = open(fmt_name_with_dir, 'w+')
            #nfd.write(out)
            #nfd.close()

            #process = subprocess.Popen("ls".split(), stdout=subprocess.PIPE)
            #out,err = process.communicate()
            #print out, err

            #renameCmd = "mv %s %s" %(fmt_name_with_dir, dir+fd.name)
            #process = subprocess.Popen(renameCmd.split(), stdout=subprocess.PIPE)
            #out,err = process.communicate()
            #print out, err


        return out

def executeCopyCommand (name, dstPath) :
    dir = dstPath
    if not os.path.exists(dir):
      os.makedirs(dir)

    copyCmd = "cp %s %s" %(name, dstPath,)
    process = subprocess.Popen(copyCmd.split(), stdout=subprocess.PIPE)
    out,err = process.communicate()

    #print out, err

    return out

def executeLocalCleanup():

    for name in os.listdir(CODE_GENERATION_PATH):
        if name.endswith(".go") or name.endswith(".thrift"):
            cmd = "rm %s" %(CODE_GENERATION_PATH+name,)
            process = subprocess.Popen(cmd.split(), stdout=subprocess.PIPE)
            out,err = process.communicate()
            #print out, err


def scan_dir_for_go_files(dir):
    for name in os.listdir(dir):
        #print "x", dir, name
        path = os.path.join(dir, name)
        if name.endswith('.go'):
            if os.path.isfile(path) and "_enum" not in path and "_func" not in path and "_db" not in path:
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

    # lets determine from the json file the structs and associated listeners
    deamons = get_listeners_from_json(goStructToListersDict)

    #pprint.pprint(goStructToListersDict)

    allCrudStructList = []
    # lets create the clientIf and .thrift files for each listener deamon
    for d in deamons:
        clientIfName = "gen" + d + "clientif.go"
        clientIfFd = open(CLIENTIF_CODE_GENERATION_PATH + clientIfName, 'w+')
        clientIfFd.write("package main\n")
        thriftFileName = d + ".thrift"
        thriftfd = open(THRIFT_CODE_GENERATION_PATH + thriftFileName, 'w+')
        thriftfd.write("namespace go %sServices\n" %(d))
        thriftfd.write("""typedef i32 int
typedef i16 uint16
""")

        # create the thrift file info
        (goMemberTypeDict, crudStructsList, goStructDict) = generate_thirft_structs_and_func(thriftfd, d, goStructToListersDict)
        thriftfd.close()

        allCrudStructList += list(set(crudStructsList).difference(set(allCrudStructList)))

        # create a client if info
        generate_clientif(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict)

    # create teh object map file
    generate_objmap(allCrudStructList)


def get_listeners_from_json(goStructToListersDict):
    deamons = []
    for dir, jsonfilename in scan_dir_for_json_files(JSON_MODEL_REGISTRAION_PATH):
        path = os.path.join(dir, jsonfilename)
        if jsonfilename.endswith(".json"):
            #print path
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
                        for d in v["Listeners"]:
                            if d not in deamons:
                                deamons.append(d)
    return deamons

def generate_thirft_structs_and_func(thriftfd, d, goStructToListersDict):
    goMemberTypeDict = {}
    goStructDict = {}
    crudStructsList = []
    for dir, gofilename in scan_dir_for_go_files(GO_MODEL_BASE_PATH):
        #print dir, gofilename, dir.split('/')[-1]

        path = os.path.join(dir, gofilename)
        gofd = open(path, 'r')
        deletingComment = False
        writingStruct = False
        memberCnt = 1
        currentStruct = None
        for line in gofd.readlines():
            if not deletingComment:
                if "struct" in line:
                    lineSplit = line.split(" ")
                    structLine = "struct " + lineSplit[1] + "{\n"
                    if lineSplit[1] in goStructToListersDict:
                        # print "found line now checking deamon", d, goStructToListersDict[lineSplit[1]]
                        if d in goStructToListersDict[lineSplit[1]]:
                            goMemberTypeDict[lineSplit[1]] = {}
                            goStructDict[lineSplit[1]] = {}
                            currentStruct = lineSplit[1]
                            thriftfd.write(structLine)
                            crudStructsList.append(lineSplit[1])
                            writingStruct = True
                elif "}" in line and writingStruct:
                    thriftfd.write("}\n")
                    if 'State' in currentStruct or "Counter" in currentStruct:
                        thriftfd.write("""struct %sGetInfo {\n\t1: int StartIdx\n\t2: int EndIdx\n\t3: int Count\n\t4: bool More\n\t5: list<%s> %sList\n}\n""" %(currentStruct, currentStruct, currentStruct))
                    writingStruct = False
                    memberCnt = 1
                # lets skip all blank lines
                # skip comments
                elif line == '\n' or \
                    "//" in line or \
                    "#" in line or \
                    "package" in line or \
                    "BaseObj" in line:
                    continue
                elif "/*" in line:
                    deletingComment = True
                elif writingStruct:  # found element in struct
                    # print "found element line", line
                    lineSplit = line.split(' ')
                    # print lineSplit
                    elemtype = lineSplit[-3].rstrip('\n') if 'KEY' in lineSplit[-1] else lineSplit[-1].rstrip('\n')

                    #print "elemtype:", lineSplit, elemtype
                    if elemtype.startswith("[]"):
                        elemtype = elemtype.lstrip("[]")
                        # lets make all list an unordered list
                        nativetype = "set<" + goToThirftTypeMap[elemtype]["native_type"] + ">"
                        goMemberTypeDict[currentStruct].update({lineSplit[0].lstrip(' ').rstrip(' ').lstrip('\t'):
                                                                nativetype})
                        goStructDict[currentStruct].update({lineSplit[0].lstrip(' ').rstrip(' ').lstrip('\t') :
                                                            elemtype + "[]"})

                        thriftfd.write("\t%s : %s %s\n" % (memberCnt,
                                                           nativetype,
                                                           lineSplit[0]))
                    else:
                        if elemtype in goToThirftTypeMap.keys():
                            goMemberTypeDict[currentStruct].update({lineSplit[0].lstrip(' ').rstrip(' ').lstrip('\t'):
                                                                        goToThirftTypeMap[elemtype]["native_type"]})
                            goStructDict[currentStruct].update({lineSplit[0].lstrip(' ').rstrip(' ').lstrip('\t') :
                                                                elemtype})

                            thriftfd.write("\t%s : %s %s\n" % (memberCnt,
                                                               goToThirftTypeMap[elemtype]["native_type"],
                                                               lineSplit[0]))
                    memberCnt += 1
            else:
                if "*/" in line:
                    deletingComment = False

    #print crudStructsList
    thriftfd.write("service %sServices {\n" % (d.upper()))
    for s in crudStructsList:
        if 'Config' in s:
            thriftfd.write(
                """\tbool Create%s(1: %s config);\n\tbool Update%s(1: %s origconfig, 2: %s newconfig, 3: list<byte> attrset);\n\tbool Delete%s(1: %s config);\n\n""" % (s, s, s, s, s, s, s))
        else: # read only objects Counters/State
            thriftfd.write("""\t%sGetInfo GetBulk%s(1: int fromIndex, 2: int count);\n""" %(s, s))
    thriftfd.write("}")
    return goMemberTypeDict, crudStructsList, goStructDict


def createClientIfCreateObject(clientIfFd, d, crudStructsList, goMemberTypeDict):
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    servicesName = d + "Services"
    clientIfFd.write("""func (clnt *%sClient) CreateObject(obj models.ConfigObj, dbHdl *sql.DB) (int64, bool) {
                        var objId int64
	                    switch obj.(type) {\n""" % (newDeamonName,))
    for s in crudStructsList:
        if 'Config' in s:
            clientIfFd.write("""
                                case models.%s :
                                data := obj.(models.%s)
                                conf := %s.New%s()\n""" % (s, s, servicesName, s))
            for k, v in goMemberTypeDict[s].iteritems():
                #print k.split(' ')
                cast = v
                # lets convert thrift i8, i16, i32, i64 to int...
                if cast.startswith('i'):
                    cast = 'int' + cast.lstrip('i')
                clientIfFd.write("""conf.%s = %s(data.%s)\n""" % (k, cast, k))
            clientIfFd.write("""
                                _, err := clnt.ClientHdl.Create%s(conf)
                                if err != nil {
                                return int64(0), false
                                }
                                objId, _ = data.StoreObjectInDb(dbHdl)
                                break\n""" % (s,))
    clientIfFd.write("""default:
		                break
	                    }

	                    return objId, true
                        }\n""")

def createClientIfDeleteObject(clientIfFd, d, crudStructsList, goMemberTypeDict):
    newDeamonName = d.upper()
    servicesName = d + "Services"
    clientIfFd.write("""func (clnt *%sClient) DeleteObject(obj models.ConfigObj, objKey string, dbHdl *sql.DB) bool {

	                    switch obj.(type) {\n""" % (newDeamonName,))
    for s in crudStructsList:
        if 'Config' in s:
            clientIfFd.write("""
                                case models.%s :
                                data := obj.(models.%s)
                                conf := %s.New%s()\n""" % (s, s, servicesName, s))
            for k, v in goMemberTypeDict[s].iteritems():
                #print k.split(' ')
                cast = v
                # lets convert thrift i8, i16, i32, i64 to int...
                if cast.startswith('i'):
                    cast = 'int' + cast.lstrip('i')
                clientIfFd.write("""conf.%s = %s(data.%s)\n""" % (k, cast, k))
            clientIfFd.write("""
                                _, err := clnt.ClientHdl.Delete%s(conf)
                                if err != nil {
                                return false
                                }
                                data.DeleteObjectFromDb(objKey, dbHdl)
                                break\n""" % (s,))
    clientIfFd.write("""default:
		                break
	                    }

	                    return true
                        }\n""")

def createClientIfGetBulkObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict):
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    servicesName = d + "Services"
    clientIfFd.write("""func (clnt *%sClient) GetBulkObject(obj models.ConfigObj, currMarker int64, count int64) (err error,
	                                objCount int64,
	                                nextMarker int64,
	                                more bool,
	                                objs []models.ConfigObj) {

	logger.Println("### Get Bulk request called with", currMarker, count)
	switch obj.(type) {
    \n""" %(newDeamonName))
    for s in crudStructsList:
        if 'State' in s or 'Counters' in s:

            clientIfFd.write("""\ncase models.%s :\n""" % (s,))

            clientIfFd.write("""
                if clnt.ClientHdl != nil {
                	var ret_obj models.%s
                    bulkInfo, _ := clnt.ClientHdl.GetBulk%s(%s.Int(currMarker), %s.Int(count))
                    if bulkInfo.Count != 0 {
                        objCount = int64(bulkInfo.Count)
                        more = bool(bulkInfo.More)
                        nextMarker = int64(bulkInfo.EndIdx)
                        for i := 0; i < int(bulkInfo.Count); i++ {
                            if len(objs) == 0 {
                                objs = make([]models.ConfigObj, 0)
                            }""" %(s, s, servicesName, servicesName))
            for k, v in goStructDict[s].iteritems():
                clientIfFd.write("""
                            ret_obj.%s = %s(bulkInfo.%sList[i].%s)""" %(k, v, s, k))
            clientIfFd.write("""\nobjs = append(objs, ret_obj)
			            }
			}
		}
		break\n""")
    clientIfFd.write("""\ndefault:
		                break
	                    }
                return nil, objCount, nextMarker, more, objs

            }\n""")

def createClientIfUpdateObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict):
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    servicesName = d + "Services"

    clientIfFd.write("""func (clnt *%sClient) UpdateObject(dbObj models.ConfigObj, obj models.ConfigObj, attrSet []byte, objKey string, dbHdl *sql.DB) bool {

	logger.Println("### Update Object called %s", attrSet, objKey)
	ok := false
	switch obj.(type) {
    """ %(newDeamonName, newDeamonName))
    for s in crudStructsList:
        if 'State' not in s and 'Counters' not in s:

            clientIfFd.write("""\ncase models.%s :""" % (s,))
            clientIfFd.write("""\n// cast original object
            origdata := dbObj.(models.%s)
            updatedata := obj.(models.%s)\n""" %(s, s) )
            clientIfFd.write("""// create new thrift objects
            origconf := %s.New%s()\nupdateconf := %s.New%s()\n""" %(servicesName, s, servicesName, s))
            for i in range(2):
                clientIfFd.write("\n")
                for k, v in goMemberTypeDict[s].iteritems():
                    cast = v
                    # lets convert thrift i8, i16, i32, i64 to int...
                    if cast.startswith('i'):
                        cast = 'int' + cast.lstrip('i')
                    if i == 0:
                        clientIfFd.write("""origconf.%s = %s(origdata.%s)\n""" % (k, cast, k))
                    else:
                        clientIfFd.write("""updateconf.%s = %s(updatedata.%s)\n""" % (k, cast, k))

            clientIfFd.write("""\n//convert attrSet to uint8 list
            newattrset := make([]int8, len(attrSet))
            for i, v := range(attrSet) {
                newattrset[i] = int8(v)
            }""")

            clientIfFd.write("""
                if clnt.ClientHdl != nil {
                    ok, err := clnt.ClientHdl.Update%s(origconf, updateconf, newattrset)
                    if ok {
                        updatedata.UpdateObjectInDb(dbObj, attrSet, dbHdl)
                    } else {
                        panic(err)
                    }
                }
                break\n""" %(s))

    clientIfFd.write("""\ndefault:
		                break
	                    }
                return ok

            }\n""")




def generate_clientif(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict):
    #newDeamonName = d[0].upper() + d[1:-1] + d[-1].upper()
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    # BELOW CODE WILL BE FORMATED BY GOFMT
    clientIfFd.write("""import (
    "%sServices"
    "models"
    "database/sql"
    )\n""" % lowerDeamonName)
    clientIfFd.write("""type %sClient struct {
	                        IPCClientBase
	                        ClientHdl *%sServices.%sServicesClient
                            }\n""" % (newDeamonName, lowerDeamonName, newDeamonName))
    clientIfFd.write("""
                        func (clnt *%sClient) Initialize(name string, address string) {
	                    clnt.Address = address
	                    return
                        }\n""" % (newDeamonName,))
    clientIfFd.write("""func (clnt *%sClient) ConnectToServer() bool {

	                    clnt.Transport, clnt.PtrProtocolFactory = CreateIPCHandles(clnt.Address)
	                    if clnt.Transport != nil && clnt.PtrProtocolFactory != nil {
		                clnt.ClientHdl = %sServices.New%sServicesClientFactory(clnt.Transport, clnt.PtrProtocolFactory)
	                    }
	                    return true
                        }\n""" % (newDeamonName, lowerDeamonName, newDeamonName))
    clientIfFd.write("""func (clnt *%sClient) IsConnectedToServer() bool {
	                    return true
                        }\n""" % (newDeamonName,))

    createClientIfCreateObject(clientIfFd, d, crudStructsList, goMemberTypeDict)
    createClientIfDeleteObject(clientIfFd, d, crudStructsList, goMemberTypeDict)
    createClientIfGetBulkObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict)
    createClientIfUpdateObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict)
    clientIfFd.close()

    # lets beautify the the client if code
    executeGoFmtCommand(clientIfFd, ["gofmt -w %s" %(clientIfFd.name,)], CLIENTIF_CODE_GENERATION_PATH)

def generate_objmap(allStructList):
    #print allStructList
    fd = open(OBJMAP_CODE_GENERATION_PATH + OBJECT_MAP_NAME, 'w+')
    fd.write("""package %s\n\n""" % MODEL_NAME)
    fd.write("""var ConfigObjectMap = map[string] ConfigObj{\n""")

    # lets temporarily add the manual objects
    fd.write(""" "IPV4Route":    &IPV4Route{},    // manually merged from originional
	"Vlan":         &Vlan{},         // manually added, no YANG defined
	"IPv4Intf":     &IPv4Intf{},     // manually added, no YANG defined
	"IPv4Neighbor": &IPv4Neighbor{}, // manually added, no YANG defined
	"BGPGlobalConfig": &BGPGlobalConfig{}, //manually added, no YANG defined
	"BGPNeighborConfig" : &BGPNeighborConfig{}, //manually added, no YANG defined\n""")

    length = len(allStructList)
    for i, s in enumerate(allStructList):
        fd.write(""""%s" : &%s{},\n""" %(s, s,))

    fd.write("""}\n""")
    fd.close()

    executeGoFmtCommand(fd, ["gofmt -w %s" %(fd.name,)], OBJMAP_CODE_GENERATION_PATH)


if __name__ == "__main__":

    for dirpath in [CODE_GENERATION_PATH,
                    CLIENTIF_CODE_GENERATION_PATH,
                    OBJMAP_CODE_GENERATION_PATH,
                    THRIFT_CODE_GENERATION_PATH]:
        if not os.path.exists(dirpath):
            os.makedirs(dirpath)

    build_thrift_from_go()
    #executeLocalCleanup()
