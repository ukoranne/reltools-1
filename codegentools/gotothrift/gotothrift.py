import os
import subprocess
import json
import pprint

OBJECT_MAP_NAME = "objectmap.go"

MODEL_NAME = 'models'
srBase = os.environ.get('SR_CODE_BASE', None)
GO_MODEL_BASE_PATH_LIST = [srBase + "/generated/src/%s/" % MODEL_NAME,
                           srBase + "/snaproute/src/models/"]
JSON_MODEL_REGISTRAION_PATH = srBase + "/snaproute/src/models/"
#JSON_MODEL_REGISTRAION_PATH = HOME + "/git/reltools/codegentools/gotojson/"
CODE_GENERATION_PATH = srBase + "/reltools/codegentools/gotothrift/"
CLIENTIF_CODE_GENERATION_PATH = srBase + "/generated/src/config/"
OBJMAP_CODE_GENERATION_PATH = srBase + "/generated/src/%s/" % MODEL_NAME
THRIFT_CODE_GENERATION_PATH = srBase + "/generated/src/gorpc/"
DBUTIL_CODE_GENERATION_PATH = THRIFT_CODE_GENERATION_PATH + "dbutils/"

daemonThriftNameChangeDict = {
    "arpd" : "arpd",
    "asicd" : "asicdServices",
    "bgpd"  : "bgpd",
    "lacpd"  : "lacpd",
    "portd" : "portdServices",
    "dhcprelayd" : "dhcprelayd",
    "stpd" : "stpd",
    "bfdd" : "bfdd"
}


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


def scan_dir_for_go_files(dirList):
    for dir in dirList:
        for name in os.listdir(dir):
            #print "x", dir, name
            path = os.path.join(dir, name)
            if name.endswith('.go'):
                if os.path.isfile(path) and "_enum" not in path and "_func" not in path and "_db" not in path:
                    yield (dir, name)
            # dir
            #elif not "." in name:
            elif os.path.isdir(path):
                for d, f  in scan_dir_for_go_files([path]):
                    yield (d, f)

def scan_dir_for_json_files(dir):

    for name in os.listdir(dir):
        #print "x", dir, name
        path = os.path.join(dir, name)
        if name.endswith('.json'):
            if os.path.isfile(path):
                yield (dir, name)
        #elif not "." in name:
        elif os.path.isdir(path):
            for d, f  in scan_dir_for_go_files([path]):
                yield (d, f)

def build_thrift_from_go():
    # generate thrift code from go code
    goStructToListersDict = {}

    # lets determine from the json file the structs and associated listeners
    deamons, accessDict = get_listeners_and_access_from_json(goStructToListersDict)

    #pprint.pprint(goStructToListersDict)

    allCrudStructList = []
    # lets create the clientIf and .thrift files for each listener deamon
    for d in deamons:
        clientIfName = "gen" + d + "clientif.go"
        clientIfFd = open(CLIENTIF_CODE_GENERATION_PATH + clientIfName, 'w+')
        clientIfFd.write("package main\n")
        thriftFileName = d + ".thrift"
        servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d
        thriftfd = open(THRIFT_CODE_GENERATION_PATH + thriftFileName, 'w+')
        thriftfd.write("namespace go %s\n" %(servicesName))
        thriftfd.write("""typedef i32 int
typedef i16 uint16
""")

        # create the thrift file info
        (goMemberTypeDict, crudStructsList, goStructDict) = generate_thirft_structs_and_func(thriftfd, d, goStructToListersDict, accessDict)
        thriftfd.close()

        allCrudStructList += list(set(crudStructsList).difference(set(allCrudStructList)))

        createConvertObjToThriftObj(d, crudStructsList, goMemberTypeDict, goStructDict, accessDict)

        # create a client if info
        generate_clientif(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict, accessDict)

    # create teh object map file
    generate_objmap(allCrudStructList)


def get_listeners_and_access_from_json(goStructToListersDict):
    deamons = []
    accessDict = {}
    for dir, jsonfilename in scan_dir_for_json_files(JSON_MODEL_REGISTRAION_PATH):
        path = os.path.join(dir, jsonfilename)
        if jsonfilename.endswith(".json"):
            #print path
            try:
                with open(path, 'r') as f:
                    data = json.load(f)

                    for k, v in data.iteritems():
                        accessDict[k] = v["access"]
                        if v["Owner"]:
                            goStructToListersDict.setdefault(k, [])
                            goStructToListersDict[k].append(v["Owner"])
                            if v["Owner"] not in deamons:
                                deamons.append(v["Owner"])
                        '''
                        NOTE: as of 1/18/2016 not being used
                        if v["Listeners"]:
                            goStructToListersDict.setdefault(k, [])
                            goStructToListersDict[k] += v["Listeners"]
                            for d in v["Listeners"]:
                                if d not in deamons:
                                    deamons.append(d)
                        '''
            except:
                pass
    return deamons, accessDict

def generate_thirft_structs_and_func(thriftfd, d, goStructToListersDict, accessDict):
    goMemberTypeDict = {}
    goStructDict = {}
    crudStructsList = []
    servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d
    for dir, gofilename in scan_dir_for_go_files(GO_MODEL_BASE_PATH_LIST):
        #print dir, gofilename, dir.split('/')[-1]

        path = os.path.join(dir, gofilename)
        gofd = open(path, 'r')
        deletingComment = False
        writingStruct = False
        memberCnt = 1
        currentStruct = None
        for line in gofd.readlines():
            if not deletingComment:
                if "//" in line:
                    line = line.split("//")[0]
                if len(line) == 0:
                    continue

                if "struct" in line:
                    lineSplit = line.split(" ")
                    structLine = "struct " + lineSplit[1] + "{\n"
                    if lineSplit[1] in goStructToListersDict:
                        # print "found line now checking deamon", d, goStructToListersDict[lineSplit[1]]
                        if d in goStructToListersDict[lineSplit[1]] and lineSplit[1] not in crudStructsList:
                            goMemberTypeDict[lineSplit[1]] = {}
                            goStructDict[lineSplit[1]] = {}
                            currentStruct = lineSplit[1]
                            thriftfd.write(structLine)
                            crudStructsList.append(lineSplit[1])
                            writingStruct = True
                elif "}" in line and writingStruct:
                    thriftfd.write("}\n")
                    if currentStruct in accessDict and 'r' in accessDict[currentStruct]:
                        thriftfd.write("""struct %sGetInfo {\n\t1: int StartIdx\n\t2: int EndIdx\n\t3: int Count\n\t4: bool More\n\t5: list<%s> %sList\n}\n""" %(currentStruct, currentStruct, currentStruct))
                    writingStruct = False
                    memberCnt = 1
                # lets skip all blank lines
                # skip comments
                elif line == '\n' or \
                    "#" in line or \
                    "package" in line or \
                    "ConfigObj" in line or \
                    ("/*" in line and "*/" in line):
                    continue
                elif "/*" in line:
                    deletingComment = True
                elif writingStruct:  # found element in struct
                    # print "found element line", line
                    lineSplit = [ x for x in line.split(' ') if x != '']
		    if len(lineSplit) == 0:
                        continue

                    #print line, lineSplit
                    elemtype = lineSplit[-3].rstrip('\n') if 'KEY' in lineSplit[-1] else lineSplit[-1].rstrip('\n')

                    #print "elemtype:", lineSplit, elemtype
                    if elemtype.startswith("[]"):
                        elemtype = elemtype.lstrip("[]")
                        # lets make all list an unordered list
                        nativetype = "list<" + goToThirftTypeMap[elemtype]["native_type"] + ">"
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
        if s in accessDict and 'w' in accessDict[s]:
            thriftfd.write(
                """\tbool Create%s(1: %s config);\n\tbool Update%s(1: %s origconfig, 2: %s newconfig, 3: list<bool> attrset);\n\tbool Delete%s(1: %s config);\n\n""" % (s, s, s, s, s, s, s))
        else: # read only objects Counters/State
            thriftfd.write("""\t%sGetInfo GetBulk%s(1: int fromIndex, 2: int count);\n""" %(s, s))
    thriftfd.write("}")
    return goMemberTypeDict, crudStructsList, goStructDict


def createClientIfCreateObject(clientIfFd, d, crudStructsList, goMemberTypeDict, accessDict):
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d
    clientIfFd.write("""func (clnt *%sClient) CreateObject(obj models.ConfigObj, dbHdl *sql.DB) (int64, bool) {
                        var objId int64
                        var err error
	                    switch obj.(type) {\n""" % (newDeamonName,))
    for s in crudStructsList:
        if s in accessDict and 'w' in accessDict[s]:
            clientIfFd.write("""
                                case models.%s :
                                data := obj.(models.%s)
                                conf := %s.New%s()\n""" % (s, s, servicesName, s))
            clientIfFd.write("""models.Convert%s%sObjToThrift(&data, conf)""" %(d, s))
            '''
            for k, v in goMemberTypeDict[s].iteritems():
                #print k.split(' ')
                cast = v
                # lets convert thrift i8, i16, i32, i64 to int...
                if cast.startswith('i'):
                    cast = 'int' + cast.lstrip('i')
                clientIfFd.write("""conf.%s = %s(data.%s)\n""" % (k, cast, k))
            '''
            clientIfFd.write("""
                                _, err = clnt.ClientHdl.Create%s(conf)
                                if err != nil {
                                fmt.Println("Create failed:", err)
                                return int64(0), false
                                }
                                objId, err = data.StoreObjectInDb(dbHdl)
                                if err != nil {
                                fmt.Println("Store in DB failed:", err)
                                return objId, false
                                }
                                break\n""" % (s,))
    clientIfFd.write("""default:
		                break
	                    }

	                    return objId, true
                        }\n""")

def createClientIfDeleteObject(clientIfFd, d, crudStructsList, goMemberTypeDict, accessDict):
    newDeamonName = d.upper()
    servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d
    clientIfFd.write("""func (clnt *%sClient) DeleteObject(obj models.ConfigObj, objKey string, dbHdl *sql.DB) bool {

	                    switch obj.(type) {\n""" % (newDeamonName,))
    for s in crudStructsList:
        if s in accessDict and 'w' in accessDict[s]:
            clientIfFd.write("""
                                case models.%s :
                                data := obj.(models.%s)
                                conf := %s.New%s()\n""" % (s, s, servicesName, s))
            clientIfFd.write("""models.Convert%s%sObjToThrift(&data, conf)""" %(d, s))
            '''
            for k, v in goMemberTypeDict[s].iteritems():
                #print k.split(' ')
                cast = v
                # lets convert thrift i8, i16, i32, i64 to int...
                if cast.startswith('i'):
                    cast = 'int' + cast.lstrip('i')
                clientIfFd.write("""conf.%s = %s(data.%s)\n""" % (k, cast, k))
            '''
            clientIfFd.write("""
                                _, err := clnt.ClientHdl.Delete%s(conf)
                                if err != nil {
                                fmt.Println("Delete failed:", err)
                                return false
                                }
                                data.DeleteObjectFromDb(objKey, dbHdl)
                                break\n""" % (s,))
    clientIfFd.write("""default:
		                break
	                    }

	                    return true
                        }\n""")

def createClientIfGetBulkObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict, accessDict):
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d
    clientIfFd.write("""func (clnt *%sClient) GetBulkObject(obj models.ConfigObj, currMarker int64, count int64) (err error,
	                                objCount int64,
	                                nextMarker int64,
	                                more bool,
	                                objs []models.ConfigObj) {

	logger.Println("### Get Bulk request called with", currMarker, count)
	switch obj.(type) {
    \n""" %(newDeamonName))
    for s in crudStructsList:
        if s in accessDict and 'r' in accessDict[s]:
            clientIfFd.write("""\ncase models.%s :\n""" % (s,))

            clientIfFd.write("""
                if clnt.ClientHdl != nil {
                	var ret_obj models.%s
                    bulkInfo, err := clnt.ClientHdl.GetBulk%s(%s.Int(currMarker), %s.Int(count))
                    if bulkInfo != nil &&bulkInfo.Count != 0 {
                        objCount = int64(bulkInfo.Count)
                        more = bool(bulkInfo.More)
                        nextMarker = int64(bulkInfo.EndIdx)
                        for i := 0; i < int(bulkInfo.Count); i++ {
                            if len(objs) == 0 {
                                objs = make([]models.ConfigObj, 0)
                            }\n""" %(s, s, servicesName, servicesName))
            for k, v in goStructDict[s].iteritems():
                if "[]" in v:
                    clientIfFd.write("""\nfor _, data := range bulkInfo.%sList[i].%s {
                            ret_obj.%s = %s(data)
                            }\n""" %(s, k, k, v.rstrip('[]')))
                else:
                    clientIfFd.write("""
                            ret_obj.%s = %s(bulkInfo.%sList[i].%s)""" %(k, v, s, k))
            clientIfFd.write("""\nobjs = append(objs, ret_obj)
			            }

			} else {
			    fmt.Println(err)
			}
		}
		break\n""")
    clientIfFd.write("""\ndefault:
		                break
	                    }
                return nil, objCount, nextMarker, more, objs

            }\n""")

def createConvertObjToThriftObj(d, crudStructsList, goMemberTypeDict, goStructDict, accessDict):

    thriftdbutilFileName = d + "dbthriftutil.go"
    thriftdbutilfd = open(DBUTIL_CODE_GENERATION_PATH + thriftdbutilFileName, 'w+')

    servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d
    thriftdbutilfd.write("package models\n")
    thriftdbutilfd.write("""import (
                         //"models"
                         "%s"
                         )""" %(servicesName))


    for s in crudStructsList:
        if s in accessDict and 'w' in accessDict[s]:
            thriftdbutilfd.write("""\nfunc Convert%s%sObjToThrift(dbobj *%s, thriftobj *%s.%s) { """ %(d, s, s, servicesName, s))
            for i, (k, v) in enumerate(goMemberTypeDict[s].iteritems()):
                #print k.split(' ')
                cast = v
                # lets convert thrift i8, i16, i32, i64 to int...
                if cast.startswith("list"):
                    cast = cast[5:-1]
                    if cast.startswith('i'):
                        cast = 'int' + cast.lstrip('i')
                    if cast == "bool":
                        thriftdbutilfd.write("""\nfor _, data%s := range dbobj.%s {
                                                      thriftobj.%s = append(thriftobj.%s, %s(data%s))
                                                  }\n""" %(i, k, k, k, cast, i))
                    else:
                        thriftdbutilfd.write("""\nfor _, data%s := range dbobj.%s {
                                                      thriftobj.%s = append(thriftobj.%s, %s(data%s))
                                                  }\n""" %(i, k, k, k, cast, i))
                else:
                    if cast.startswith('i'):
                        cast = 'int' + cast.lstrip('i')

                    thriftdbutilfd.write("""thriftobj.%s = %s(dbobj.%s)\n""" % (k, cast, k))
            thriftdbutilfd.write("""}\n""")

            thriftdbutilfd.write("""\nfunc ConvertThriftTo%s%sObj(thriftobj *%s.%s, dbobj *%s) { """ %(d, s, servicesName, s, s))
            for i, (k, v) in enumerate(goStructDict[s].iteritems()):
                #print k.split(' ')
                cast = v

                # lets convert thrift i8, i16, i32, i64 to int...
                if cast.endswith("[]"):
                    cast = cast[:-2]
                    thriftdbutilfd.write("""\nfor _, data%s := range thriftobj.%s {
                                                  dbobj.%s = append(dbobj.%s, %s(data%s))
                                              }\n""" %(i, k, k, k, cast, i))
                else:
                    thriftdbutilfd.write("""dbobj.%s = %s(thriftobj.%s)\n""" % (k, cast, k))

            thriftdbutilfd.write("""}\n""")
    thriftdbutilfd.close()
    executeGoFmtCommand(thriftdbutilfd, ["gofmt -w %s" %(thriftdbutilfd.name,)], DBUTIL_CODE_GENERATION_PATH)


def createClientIfUpdateObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict, accessDict):
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d

    clientIfFd.write("""func (clnt *%sClient) UpdateObject(dbObj models.ConfigObj, obj models.ConfigObj, attrSet []bool, objKey string, dbHdl *sql.DB) bool {
        var ok bool
        var err error
	logger.Println("### Update Object called %s", attrSet, objKey)
	ok = false
        err = nil
	switch obj.(type) {
    """ %(newDeamonName, newDeamonName))
    for s in crudStructsList:
        if s in accessDict and 'w' in accessDict[s]:
            clientIfFd.write("""\ncase models.%s :""" % (s,))
            clientIfFd.write("""\n// cast original object
            origdata := dbObj.(models.%s)
            updatedata := obj.(models.%s)\n""" %(s, s) )
            clientIfFd.write("""// create new thrift objects
            origconf := %s.New%s()\nupdateconf := %s.New%s()\n""" %(servicesName, s, servicesName, s))
            #for i in range(2):
            clientIfFd.write("""models.Convert%s%sObjToThrift(&origdata, origconf)
            models.Convert%s%sObjToThrift(&updatedata, updateconf)""" %(d, s, d, s))
            '''
            for k, v in goMemberTypeDict[s].iteritems():
                cast = v
                # lets convert thrift i8, i16, i32, i64 to int...
                if cast.startswith('i'):
                    cast = 'int' + cast.lstrip('i')
                if i == 0:
                    clientIfFd.write("""origconf.%s = %s(origdata.%s)\n""" % (k, cast, k))
                else:
                    clientIfFd.write("""updateconf.%s = %s(updatedata.%s)\n""" % (k, cast, k))
            '''
            clientIfFd.write("""
                if clnt.ClientHdl != nil {
                    ok, err = clnt.ClientHdl.Update%s(origconf, updateconf, attrSet)
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




def generate_clientif(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict, accessDict):
    #newDeamonName = d[0].upper() + d[1:-1] + d[-1].upper()
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d

    print crudStructsList
    # BELOW CODE WILL BE FORMATED BY GOFMT
    clientIfFd.write("""import (
    "%s"
    "fmt"
    "models"
    "database/sql"
    "utils/ipcutils"
    )\n""" % servicesName)
    clientIfFd.write("""type %sClient struct {
	                        ipcutils.IPCClientBase
	                        ClientHdl *%s.%sServicesClient
                            }\n""" % (newDeamonName, servicesName, newDeamonName))
    clientIfFd.write("""
                        func (clnt *%sClient) Initialize(name string, address string) {
	                    clnt.Address = address
	                    return
                        }\n""" % (newDeamonName,))
    clientIfFd.write("""func (clnt *%sClient) ConnectToServer() bool {

	                    clnt.TTransport, clnt.PtrProtocolFactory, _ = ipcutils.CreateIPCHandles(clnt.Address)
	                    if clnt.TTransport != nil && clnt.PtrProtocolFactory != nil {
		                clnt.ClientHdl = %s.New%sServicesClientFactory(clnt.TTransport, clnt.PtrProtocolFactory)
		                if clnt.ClientHdl != nil {
		                    clnt.IsConnected = true
		                } else {
		                    clnt.IsConnected = false
		                }
	                    }
	                    return true
                        }\n""" % (newDeamonName, servicesName, newDeamonName))
    clientIfFd.write("""func (clnt *%sClient) IsConnectedToServer() bool {
	                    return clnt.IsConnected
                        }\n""" % (newDeamonName,))

    createClientIfCreateObject(clientIfFd, d, crudStructsList, goMemberTypeDict, accessDict)
    createClientIfDeleteObject(clientIfFd, d, crudStructsList, goMemberTypeDict, accessDict)
    createClientIfGetBulkObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict, accessDict)
    createClientIfUpdateObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict, accessDict)
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
                    THRIFT_CODE_GENERATION_PATH,
                    DBUTIL_CODE_GENERATION_PATH]:
        if not os.path.exists(dirpath):
            os.makedirs(dirpath)

    build_thrift_from_go()
    #executeLocalCleanup()
