import os
import subprocess
import json
import pprint
import re

OBJECT_MAP_NAME = "objectmap.go"

MODEL_NAME = 'models'
srBase = os.environ.get('SR_CODE_BASE', None)
GO_MODEL_BASE_PATH_LIST = [srBase + "/generated/src/%s/" % MODEL_NAME,
                           srBase + "/snaproute/src/models/"]
JSON_MODEL_REGISTRAION_PATH = srBase + "/snaproute/src/models/"
THRIFT_UTILS_PATH = srBase + "/snaproute/src/models/"
CLIENTIF_SRC_PATCH = srBase + "/snaproute/src/config/"
#JSON_MODEL_REGISTRAION_PATH = HOME + "/git/reltools/codegentools/gotojson/"
CODE_GENERATION_PATH = srBase + "/reltools/codegentools/gotothrift/"
CLIENTIF_CODE_GENERATION_PATH = srBase + "/generated/src/config/"
CLIENTIF_FILE_PATH = srBase + "/src/config/"
SRC_BASE = srBase + "/snaproute/src/"
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

class DaemonObjectsInfo (object) :
    def __init__ (self, name, location):
        self.name   =  name
        self.location =  location
        self.thriftFileName = SRC_BASE + location + '/'+  name + ".thrift"
        self.thriftUtilsFileName = THRIFT_UTILS_PATH + name + "dbthriftutil.go"
	self.clientIfFileName = CLIENTIF_SRC_PATCH + name + "clientif.go"
        self.objectDict = {}

    def __str__(self):
        print '%s' %(self.name)
        for name, obj in self.objectDict.iteritems(): 
            print '\t%s  - %s' %(name, obj)
        return ''

    def parseSrcFile(self):
        for objName, info in self.objectDict.iteritems():
            infoFile = srBase+ '/reltools/codegentools/._genInfo/'+objName + 'Members.json'
            with open(infoFile) as fileHdl:
                objMembersData = json.load(fileHdl)
            objMembersStrData = {}
            for attrName, attrInfo in objMembersData.iteritems():
                attrDict = {} 
                for key, val in attrInfo.iteritems() :
                    attrDict[str(key)] = str(val)
                objMembersStrData[str(attrName)] = attrDict

            info['membersInfo'] = objMembersStrData
            
        #print self.objectDict


    def generateThriftInterfaces(self, objectNames):
        thriftfd = open(self.thriftFileName, 'w+')
        dmn = self.name
        thriftfd.write("namespace go %s\n" %(dmn))                                                             
        thriftfd.write("""typedef i32 int\ntypedef i16 uint16\n""")    

        for structName, structInfo in objectNames.objectDict.iteritems ():
            structName = str(structName)
            srcFile = JSON_MODEL_REGISTRAION_PATH + structInfo['srcfile']
            srcFileFd = open(srcFile, 'r')
            deletingComment = False
            writingStruct = False
            memberCnt = 1
            currentStruct = None
            for line in srcFileFd.readlines():
                if not deletingComment:
                    if "//" in line:
                        line = line.split("//")[0]
                    if len(line) == 0:
                        continue

                    if "struct" in line:
                        lineSplit = line.split(" ")
                        structLine = "struct " + lineSplit[1] + "{\n"
                        currentStruct = lineSplit[1]
                        if structName in lineSplit and 'type' in lineSplit and 'struct' in lineSplit:
                            writingStruct = True
                            thriftfd.write(structLine)                                                                  
                    elif "}" in line and writingStruct:
                        thriftfd.write("}\n")
                        if structInfo['access'] == 'r' and currentStruct == structName:
                            thriftfd.write("""struct %sGetInfo {\n\t1: int StartIdx\n\t2: int EndIdx\n\t3: int Count\n\t4: bool More\n\t5: list<%s> %sList\n}\n""" %(currentStruct, currentStruct, currentStruct))
                        writingStruct = False
                        memberCnt = 1
                    elif line == '\n' or \
                        "#" in line or \
                        "package" in line or \
                        "BaseObj" in line or \
                        ("/*" in line and "*/" in line):
                        continue
                    elif "/*" in line:
                        deletingComment = True
                    elif writingStruct:  # found element in struct
                        lineSplit = [ re.sub(r'\W+', '',x) for x in line.split(' ') if x != '']
                        if 'KEY' in lineSplit:
                            elemtype = lineSplit[lineSplit.index('KEY') -2]
                        else:
                            elemtype = lineSplit[-1].rstrip('\n')
                        if elemtype.startswith("[]"):
                            elemtype = elemtype.lstrip("[]")
                            nativetype = "list<" + goToThirftTypeMap[elemtype]["native_type"] + ">"
                            thriftfd.write("\t%s : %s %s\n" % (memberCnt,
                                                            nativetype,
                                                            lineSplit[0]))
                        else:
                            if elemtype in goToThirftTypeMap.keys():
                                thriftfd.write("\t%s : %s %s\n" % (memberCnt,
                                                                goToThirftTypeMap[elemtype]["native_type"],
                                                                lineSplit[0]))
                        memberCnt += 1
                else:
                    if "*/" in line:
                        deletingComment = False

        thriftfd.write("service %sServices {\n" % (dmn.upper()))
        for structName, structInfo in objectNames.objectDict.iteritems ():
            s = structName
            if structInfo['access'] == 'w' or structInfo['access'] == 'rw':
                thriftfd.write(
                    """\tbool Create%s(1: %s config);\n\tbool Update%s(1: %s origconfig, 2: %s newconfig, 3: list<bool> attrset);\n\tbool Delete%s(1: %s config);\n\n""" % (s, s, s, s, s, s, s))
            else: # read only objects Counters/State
                thriftfd.write("""\t%sGetInfo GetBulk%s(1: int fromIndex, 2: int count);\n""" %(s, s))
        thriftfd.write("}")
        thriftfd.close()
        print 'Thrift file for %s is %s' %(dmn, self.thriftFileName)
        return 


    def createConvertObjToThriftObj(self, objectNames):
        if self.name == 'ospfd':  ### Hari TODO remove this condition once OSPF works
            return
        print '#ThriftUtils file is %s' %(self.thriftUtilsFileName)
        thriftdbutilfd = open(self.thriftUtilsFileName, 'w+')

        servicesName = self.name
        thriftdbutilfd.write("package models\n")
        thriftdbutilfd.write("""import (\n 
                                "%s"\n)""" %(servicesName))

        for structName, structInfo in objectNames.objectDict.iteritems ():
            structName = str(structName)
            s = structName
            d = self.name
            if structInfo['access'] in ['w', 'rw']:
                thriftdbutilfd.write("""\nfunc Convert%s%sObjToThrift(dbobj *%s, thriftobj *%s.%s) { """ %(d, s, s, servicesName, s))
                for i, (k, v) in enumerate(structInfo['membersInfo'].iteritems()):
                    attrType = v['type']
                    if v['isArray'] != 'False':
                         thriftdbutilfd.write("""\nfor _, data%s := range dbobj.%s {
                                                        thriftobj.%s = append(thriftobj.%s, %s(data%s))
                                                    }\n""" %(i, k, k, k, attrType, i))
                    else:
                        thriftdbutilfd.write("""thriftobj.%s = %s(dbobj.%s)\n""" % (k, attrType, k)) 
                thriftdbutilfd.write("""}\n""")


                thriftdbutilfd.write("""\nfunc ConvertThriftTo%s%sObj(thriftobj *%s.%s, dbobj *%s) { """ %(d, s, servicesName, s, s))

                for i, (k, v) in enumerate(structInfo['membersInfo'].iteritems()):
                    attrInfo = v
                    attrType = attrInfo['type']
                    if attrInfo['isArray'] != 'False':
                        thriftdbutilfd.write("""\nfor _, data%s := range thriftobj.%s {
                                                    dbobj.%s = append(dbobj.%s, %s(data%s))
                                                }\n""" %(i, k, k, k, attrType, i))
                    else:
                        thriftdbutilfd.write("""dbobj.%s = %s(thriftobj.%s)\n""" % (k, attrType, k))

                thriftdbutilfd.write("""}\n""")
        thriftdbutilfd.close()
   
    def clientIfBasicHelper(self, clientIfFd, servicesName, newDeamonName, lowerDeamonName):
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

    def createClientIfCreateObject(self, clientIfFd, servicesName, newDeamonName, lowerDeamonName, objectNames):
	print 'Create clientIf Create Object for %s' %(self.name)
	clientIfFd.write("""func (clnt *%sClient) CreateObject(obj models.ConfigObj, dbHdl *sql.DB) (int64, bool) {
			    var objId int64
				switch obj.(type) {\n""" % (newDeamonName,))
        for structName, structInfo in objectNames.objectDict.iteritems ():
            structName = str(structName)
            s = structName
            d = self.name
            if structInfo['access'] in ['w', 'rw']:
		clientIfFd.write("""
				    case models.%s :
				    data := obj.(models.%s)
				    conf := %s.New%s()\n""" % (s, s, servicesName, s))
		clientIfFd.write("""models.Convert%s%sObjToThrift(&data, conf)""" %(d, s))
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

    def generate_clientif(self, objectNames):
        if self.name == 'ospfd':  ### Hari TODO remove this condition once OSPF works
            return
	clientIfFd = open(self.clientIfFileName, 'w+')
        print '#ClientIf file is %s' %(self.clientIfFileName)
        servicesName = self.name
	newDeamonName = servicesName.upper()
	lowerDeamonName = servicesName.lower()
        #print objectNames.objectDict
	clientIfFd.write("package main\n")
	#if (len([ x for x,y in accessDict.iteritems() if x in crudStructsList and 'r' in y]) > 0):
	if (len([x for x,y in objectNames.objectDict.iteritems() if x in objectNames.objectDict and 'r' in y]) > 0):
	    # BELOW CODE WILL BE FORMATED BY GOFMT
	    clientIfFd.write("""import (
	    "%s"
	    "fmt"
	    "models"
	    "database/sql"
	    "utils/ipcutils"
	    )\n""" % servicesName)
	else:
	    # BELOW CODE WILL BE FORMATED BY GOFMT
	    clientIfFd.write("""import (
	    "%s"
	    "models"
	    "database/sql"
	    "utils/ipcutils"
	    )\n""" % servicesName)
	self.clientIfBasicHelper(clientIfFd, servicesName, newDeamonName, lowerDeamonName)
	self.createClientIfCreateObject(clientIfFd, servicesName, newDeamonName, lowerDeamonName, objectNames)
        '''
    createClientIfDeleteObject(clientIfFd, d, crudStructsList, goMemberTypeDict, accessDict)
    createClientIfGetBulkObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict, accessDict)
    createClientIfUpdateObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict, accessDict)
    clientIfFd.close()

        '''

gDryRun =  False
def generateThriftAndClientIfs():
    # generate thrift code from go code
    genObjInfoJson = JSON_MODEL_REGISTRAION_PATH + 'genObjectConfig.json'
    goDmnDirsInfoJson = JSON_MODEL_REGISTRAION_PATH + 'goObjInfo.json'
    yangDmnDirsInfoJson = JSON_MODEL_REGISTRAION_PATH + 'yangObjInfo.json'

    ownerDirsInfo = {} 
    for dirFile  in [goDmnDirsInfoJson, yangDmnDirsInfoJson]:
        with open(dirFile) as locnFile:
            objData = json.load(locnFile)

        for dmn, info in objData.iteritems():
            if not ownerDirsInfo.has_key(info['owner']):
                ownerDirsInfo[info['owner']] = info['location']

    
    with open(genObjInfoJson) as infoFile:
        objData = json.load(infoFile)

    ownerToObjMap = {}
    for name,  dtls in objData.iteritems():
        if ownerToObjMap.has_key(dtls['owner']):
            dmnObj = ownerToObjMap[dtls['owner']]
        else:
            dmnObj = DaemonObjectsInfo (dtls['owner'], ownerDirsInfo[dtls['owner']]) 
            ownerToObjMap[dtls['owner']] = dmnObj

        dmnObj.objectDict[name] = dtls
    
    for dmn, entry in ownerToObjMap.iteritems():
        clientIfFileName = CLIENTIF_FILE_PATH + dmn + '/' +"clientif.go"
        thriftFileName = SRC_BASE + ownerDirsInfo[dmn] + '/' + dmn + ".thrift"
        entry.parseSrcFile()
        entry.generateThriftInterfaces(ownerToObjMap[dmn])
        entry.createConvertObjToThriftObj(ownerToObjMap[dmn])
        entry.generate_clientif(ownerToObjMap[dmn])

    return
    goStructToListersDict = {}


    allCrudStructList = []
    # lets create the clientIf and .thrift files for each listener deamon
    #for d in deamons:
    #    createConvertObjToThriftObj(d, crudStructsList, goMemberTypeDict, goStructDict, accessDict)

        # create a client if info
        # Jay Start here
        # generate_clientif(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict, accessDict)

    # create teh object map file
    # generate_objmap(allCrudStructList)


def createClientIfCreateObject(clientIfFd, d, crudStructsList, goMemberTypeDict, accessDict):
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d
    clientIfFd.write("""func (clnt *%sClient) CreateObject(obj models.ConfigObj, dbHdl *sql.DB) (int64, bool) {
                        var objId int64
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

def createClientIfUpdateObject(clientIfFd, d, crudStructsList, goMemberTypeDict, goStructDict, accessDict):
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d

    clientIfFd.write("""func (clnt *%sClient) UpdateObject(dbObj models.ConfigObj, obj models.ConfigObj, attrSet []bool, objKey string, dbHdl *sql.DB) bool {

	logger.Println("### Update Object called %s", attrSet, objKey)
	ok := false
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
                    ok, err := clnt.ClientHdl.Update%s(origconf, updateconf, attrSet)
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
    newDeamonName = d.upper()
    lowerDeamonName = d.lower()
    servicesName = daemonThriftNameChangeDict[d] if d in daemonThriftNameChangeDict else d

    print crudStructsList
    if (len([ x for x,y in accessDict.iteritems() if x in crudStructsList and 'r' in y]) > 0):
        # BELOW CODE WILL BE FORMATED BY GOFMT
        clientIfFd.write("""import (
        "%s"
        "fmt"
        "models"
        "database/sql"
        "utils/ipcutils"
        )\n""" % servicesName)
    else:
        # BELOW CODE WILL BE FORMATED BY GOFMT
        clientIfFd.write("""import (
        "%s"
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

def generate_objmap(allStructList):
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



if __name__ == "__main__":

    for dirpath in [CODE_GENERATION_PATH,
                    CLIENTIF_CODE_GENERATION_PATH,
                    OBJMAP_CODE_GENERATION_PATH,
                    THRIFT_CODE_GENERATION_PATH,
                    DBUTIL_CODE_GENERATION_PATH]:
        if not os.path.exists(dirpath):
            os.makedirs(dirpath)

    generateThriftAndClientIfs()
