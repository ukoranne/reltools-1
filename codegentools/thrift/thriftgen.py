import os
import json
import re

OBJECT_MAP_NAME = "genObjMap.go"

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
OBJMAP_CODE_GENERATION_PATH = srBase + "/snaproute/src/%s/" % MODEL_NAME
THRIFT_CODE_GENERATION_PATH = srBase + "/generated/src/gorpc/"
DBUTIL_CODE_GENERATION_PATH = THRIFT_CODE_GENERATION_PATH + "dbutils/"
GENERATED_FILES_LIST = srBase + "/reltools/codegentools/._genInfo/generatedGoFiles.txt"

GENERATED_FILES_LISTING_FILE = srBase + '/reltools/codegentools/._genInfo'
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
    def __init__ (self, name, location, svcName):
        self.name   =  name
        self.location =  location
        self.thriftFileName = SRC_BASE + location + '/'+  name + ".thrift"
        self.thriftUtilsFileName = THRIFT_UTILS_PATH + "gen_" + name + "dbthriftutil.go"
        self.clientIfFileName = CLIENTIF_SRC_PATCH + "gen_" + name + "clientif.go"
        self.servicesName = self.name
        self.SName = svcName
        self.newDeamonName = self.servicesName.upper()
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
            

    def addGeneratedFilesNamesToListing (self):
        with open(GENERATED_FILES_LIST, 'a+') as fp:
            fp.write(self.thriftFileName+ '\n')
            fp.write(self.thriftUtilsFileName+ '\n')
            fp.write(self.clientIfFileName+ '\n')

    def generateThriftInterfaces(self, objectNames):
        thriftfd = open(self.thriftFileName, 'w+')
        sName = self.SName
        if sName != "nil" :
            thriftfd.write("include \"%s.thrift\"\n" %(sName))
        dmn = self.name
        thriftfd.write("namespace go %s\n" %(dmn))                                                             
        thriftfd.write("""typedef i32 int\ntypedef i16 uint16\n""")    
        for structName, structInfo in objectNames.objectDict.iteritems ():
            line = 'struct ' + structName + ' {'
            thriftfd.write(line + '\n')
            index = 0
            for attrName, attrInfo in structInfo['membersInfo'].iteritems():
                index = index+1
                if attrInfo['isArray'] != 'False' :
                    if str(attrInfo['type']) in goToThirftTypeMap:
                        nativetype = "list<" + goToThirftTypeMap[str(attrInfo['type'])]["native_type"] + ">"
                        thriftfd.write("\t%s : %s %s\n" % (index,
                                                           nativetype,
                                                           attrName))
                    else:
                        thriftfd.write("\t%s : %s %s\n" % (index,
                                                           "list<" + str(attrInfo['type']) + ">",
                                                           attrName))
                else:
                    if str(attrInfo['type']) in goToThirftTypeMap:
                        thriftfd.write("\t%s : %s %s\n" % (index,
                                                         goToThirftTypeMap[str(attrInfo['type'])]['native_type'],
                                                         attrName))
                    else:
                        thriftfd.write("\t%s : %s %s\n" % (index,
                                                         str(attrInfo['type']),
                                                         attrName))
             
            thriftfd.write('}\n')
            if structInfo['access'] == 'r' :
                thriftfd.write("""struct %sGetInfo {\n\t1: int StartIdx\n\t2: int EndIdx\n\t3: int Count\n\t4: bool More\n\t5: list<%s> %sList\n}\n""" %(structName, structName, structName))


        if sName == "nil" :
            thriftfd.write("service %sServices {\n" % (dmn.upper()))
        else :
            thriftfd.write("service %sServices extends %s.%sServices {\n" % (dmn.upper(), sName, sName.upper()))
        for structName, structInfo in objectNames.objectDict.iteritems ():
            s = structName
            if structInfo['access'] == 'w' or structInfo['access'] == 'rw':
                thriftfd.write(
                    """\tbool Create%s(1: %s config);\n\tbool Update%s(1: %s origconfig, 2: %s newconfig, 3: list<bool> attrset);\n\tbool Delete%s(1: %s config);\n\n""" % (s, s, s, s, s, s, s))
            if 'r' in structInfo['access']: # read only objects Counters/State
                thriftfd.write("""\t%sGetInfo GetBulk%s(1: int fromIndex, 2: int count);\n""" %(s, s))
        thriftfd.write("}")
        thriftfd.close()
        print 'Thrift file for %s is %s' %(dmn, self.thriftFileName)
        return 


    def createConvertObjToThriftObj(self, objectNames):
        thriftdbutilfd = open(self.thriftUtilsFileName, 'w+')

        thriftdbutilfd.write("package models\n")
        thriftdbutilfd.write("""import (\n 
                                "%s"\n)""" %(self.servicesName))

        for structName, structInfo in objectNames.objectDict.iteritems ():
            structName = str(structName)
            s = structName
            d = self.name
            if structInfo['access'] in ['w', 'rw', 'r', '']:
                thriftdbutilfd.write("""\nfunc Convert%s%sObjToThrift(dbobj *%s, thriftobj *%s.%s) { """ %(d, s, s, self.servicesName, s))
                for i, (k, v) in enumerate(structInfo['membersInfo'].iteritems()):
                    attrType = v['type'][1:] if v['type'].startswith('u') else v['type']

                    if v['isArray'] != 'False':
                        if attrType in goToThirftTypeMap:
                            thriftdbutilfd.write("""\nfor _, data%s := range dbobj.%s {
                                                        thriftobj.%s = append(thriftobj.%s, %s(data%s))
                                                    }\n""" %(i, k, k, k, attrType, i))
                        else:
                            thriftdbutilfd.write("""\nfor _, data%s := range dbobj.%s {
                                                        thriftdata%s := new(%s.%s)
                                                        Convert%s%sObjToThrift(&data%s, thriftdata%s)
                                                        thriftobj.%s = append(thriftobj.%s, thriftdata%s)
                                                    }\n""" %(i, k, i, self.servicesName, attrType, d, attrType, i, i, k, k, i))
                    else:
                        if attrType in goToThirftTypeMap:
                            thriftdbutilfd.write("""thriftobj.%s = %s(dbobj.%s)\n""" % (k, attrType, k)) 
                        else:
                            thriftdbutilfd.write("""Convert%s%sObjToThrift(&dbobj.%s, thriftobj.%s)\n""" % (d, attrType, k, k)) 
                thriftdbutilfd.write("""}\n""")


                thriftdbutilfd.write("""\nfunc ConvertThriftTo%s%sObj(thriftobj *%s.%s, dbobj *%s) { """ %(d, s, self.servicesName, s, s))

                for i, (k, v) in enumerate(structInfo['membersInfo'].iteritems()):
                    attrInfo = v
                    #attrType = v['type'][1:] if v['type'].startswith('u') else v['type']
                    attrType = attrInfo['type']
                    if attrInfo['isArray'] != 'False':
                        if attrType in goToThirftTypeMap:
                            thriftdbutilfd.write("""\nfor _, data%s := range thriftobj.%s {
                                                    dbobj.%s = append(dbobj.%s, %s(data%s))
                                                }\n""" %(i, k, k, k, attrType, i))
                        else:
                            thriftdbutilfd.write("""\nfor _, thriftdata%s := range thriftobj.%s {
                                                        dbobjdata%s := new(%s)
                                                        ConvertThriftTo%s%sObj(thriftdata%s, dbobjdata%s)
                                                        dbobj.%s = append(dbobj.%s, *dbobjdata%s)
                                                    }\n""" %(i, k, i, attrType, d, attrType, i, i, k, k, i))
                    else:
                        if attrType in goToThirftTypeMap:
                            thriftdbutilfd.write("""dbobj.%s = %s(thriftobj.%s)\n""" % (k, attrType, k))
                        else:
                            thriftdbutilfd.write("""ConvertThriftTo%s%sObj(thriftobj.%s, &dbobj.%s)\n""" % (d, attrType, k, k)) 

                thriftdbutilfd.write("""}\n""")
        thriftdbutilfd.close()
   
    def clientIfBasicHelper(self, clientIfFd):
        clientIfFd.write("""type %sClient struct {
                                    ipcutils.IPCClientBase
                                    ClientHdl *%s.%sServicesClient
                                }\n""" % (self.newDeamonName, self.servicesName, self.newDeamonName))
        clientIfFd.write("""
                            func (clnt *%sClient) Initialize(name string, address string) {
                                clnt.Address = address
                                return
                            }\n""" % (self.newDeamonName,))
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
                            }\n""" % (self.newDeamonName, self.servicesName, self.newDeamonName))
        clientIfFd.write("""func (clnt *%sClient) IsConnectedToServer() bool {
                                return clnt.IsConnected
                            }\n""" % (self.newDeamonName,))

    def createClientIfCreateObject(self, clientIfFd, objectNames):
        clientIfFd.write("""func (clnt *%sClient) CreateObject(obj models.ConfigObj, dbHdl *sql.DB) (int64, bool) {
                            var objId int64
                            var err error
                                switch obj.(type) {\n""" % (self.newDeamonName,))
        for structName, structInfo in objectNames.objectDict.iteritems ():
            structName = str(structName)
            s = structName
            d = self.name
            if structInfo['access'] in ['w', 'rw']:
                clientIfFd.write("""
                                    case models.%s :
                                    data := obj.(models.%s)
                                    conf := %s.New%s()\n""" % (s, s, self.servicesName, s))
                clientIfFd.write("""models.Convert%s%sObjToThrift(&data, conf)""" %(d, s))
                clientIfFd.write("""
                                    _, err = clnt.ClientHdl.Create%s(conf)
                                    if err != nil {
								fmt.Println("Create failed:", err)
                                    return int64(0), false
                                    }
                                    objId, err = data.StoreObjectInDb(dbHdl)
                                    if err != nil {
								fmt.Println("Store object in DB failed:", err)
                                    return objId, false
                                    }
                                    break\n""" % (s,))
        clientIfFd.write("""default:
                                    break
                                }

                                return objId, true
                            }\n""")

    def createClientIfDeleteObject(self, clientIfFd, objectNames):
        clientIfFd.write("""func (clnt *%sClient) DeleteObject(obj models.ConfigObj, objKey string, dbHdl *sql.DB) bool {

                                switch obj.(type) {\n""" % (self.newDeamonName,))
        for structName, structInfo in objectNames.objectDict.iteritems ():
            structName = str(structName)
            s = structName
            d = self.name
            if structInfo['access'] in ['w', 'rw']:
                clientIfFd.write("""
                                    case models.%s :
                                    data := obj.(models.%s)
                                    conf := %s.New%s()\n""" % (s, s, self.servicesName, s))
                clientIfFd.write("""models.Convert%s%sObjToThrift(&data, conf)""" %(d, s))
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

    def createClientIfGetObject(self, clientIfFd, objectNames):
        clientIfFd.write("""func (clnt *%sClient) GetObject(obj models.ConfigObj) (models.ConfigObj, bool) {\n""" % (self.newDeamonName,))
        clientIfFd.write("""return nil, false                                                                                                    
                            }\n""")

    def createClientIfUpdateObject(self, clientIfFd, objectNames):
        clientIfFd.write("""func (clnt *%sClient) UpdateObject(dbObj models.ConfigObj, obj models.ConfigObj, attrSet []bool, objKey string, dbHdl *sql.DB) bool {

            var ok bool
            var err error
	    logger.Println("### Update Object called %s", attrSet, objKey)
	    ok = false
            err = nil
            switch obj.(type) {
        """ %(self.newDeamonName, self.newDeamonName))
        for structName, structInfo in objectNames.objectDict.iteritems ():
            structName = str(structName)
            s = structName
            d = self.name
            if structInfo['access'] in ['w', 'rw']:
                clientIfFd.write("""\ncase models.%s :""" % (s,))
                clientIfFd.write("""\n// cast original object
                origdata := dbObj.(models.%s)
                updatedata := obj.(models.%s)\n""" %(s, s) )
                clientIfFd.write("""// create new thrift objects
                origconf := %s.New%s()\nupdateconf := %s.New%s()\n""" %(self.servicesName, s, self.servicesName, s))
                clientIfFd.write("""models.Convert%s%sObjToThrift(&origdata, origconf)
                models.Convert%s%sObjToThrift(&updatedata, updateconf)""" %(d, s, d, s))
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

    def createClientIfGetBulkObject(self, clientIfFd, objectNames):
        clientIfFd.write("""func (clnt *%sClient) GetBulkObject(obj models.ConfigObj, currMarker int64, count int64) (err error,
                                            objCount int64,
                                            nextMarker int64,
                                            more bool,
                                            objs []models.ConfigObj) {

            logger.Println("### Get Bulk request called with", currMarker, count)
            switch obj.(type) {
        \n""" %(self.newDeamonName))
        for structName, structInfo in objectNames.objectDict.iteritems ():
            structName = str(structName)
            s = structName
            d = self.name
            if 'r' in structInfo['access']:
                clientIfFd.write("""\ncase models.%s :\n""" % (s,))

                clientIfFd.write("""
                    if clnt.ClientHdl != nil {
                        bulkInfo, err := clnt.ClientHdl.GetBulk%s(%s.Int(currMarker), %s.Int(count))
                        if bulkInfo != nil &&bulkInfo.Count != 0 {
                            objCount = int64(bulkInfo.Count)
                            more = bool(bulkInfo.More)
                            nextMarker = int64(bulkInfo.EndIdx)
                            for i := 0; i < int(bulkInfo.Count); i++ {
                                ret_obj := new(models.%s)
                                if len(objs) == 0 {
                                    objs = make([]models.ConfigObj, 0)
                                }\n""" %(s, self.servicesName, self.servicesName, s))

                clientIfFd.write("""\nmodels.ConvertThriftTo%s%sObj(bulkInfo.%sList[i], ret_obj)""" % (d, s, s))
                clientIfFd.write("""\nobjs = append(objs, ret_obj)
                                        }

                            } else {
                                logger.Println(err)
                            }
                    }
                    break\n""")
        clientIfFd.write("""\ndefault:
                                    break
                                }
                    return nil, objCount, nextMarker, more, objs

                }\n""")

    def generate_clientif(self, objectNames):
        clientIfFd = open(self.clientIfFileName, 'w+')
        clientIfFd.write("package main\n")
        #if (len([ x for x,y in accessDict.iteritems() if x in crudStructsList and 'r' in y]) > 0):
        # BELOW CODE WILL BE FORMATED BY GOFMT
        clientIfFd.write("""import (\n "%s"\n"fmt"\n"models"\n"database/sql"\n"utils/ipcutils")\n""" % self.servicesName)
        self.clientIfBasicHelper(clientIfFd)
        self.createClientIfCreateObject(clientIfFd, objectNames)
        self.createClientIfDeleteObject(clientIfFd, objectNames)
        self.createClientIfUpdateObject(clientIfFd, objectNames)
        self.createClientIfGetBulkObject(clientIfFd, objectNames)
        self.createClientIfGetObject(clientIfFd, objectNames)
        
        clientIfFd.close()

gDryRun =  False
def generateThriftAndClientIfs():
    # generate thrift code from go code
    genObjInfoJson = JSON_MODEL_REGISTRAION_PATH + 'genObjectConfig.json'
    goDmnDirsInfoJson = JSON_MODEL_REGISTRAION_PATH + 'goObjInfo.json'
    yangDmnDirsInfoJson = JSON_MODEL_REGISTRAION_PATH + 'yangObjInfo.json'

    ownerDirsInfo = {} 
    ownerInternalServiceInfo = {}
    for dirFile  in [goDmnDirsInfoJson, yangDmnDirsInfoJson]:
        with open(dirFile) as locnFile:
            objData = json.load(locnFile)

        for dmn, info in objData.iteritems():
            if not ownerDirsInfo.has_key(info['owner']):
                ownerDirsInfo[info['owner']] = info['location']
                ownerInternalServiceInfo[info['owner']] = info['svcName']

    
    with open(genObjInfoJson) as infoFile:
        objData = json.load(infoFile)

    ownerToObjMap = {}
    for name,  dtls in objData.iteritems():
        if ownerToObjMap.has_key(dtls['owner']):
            dmnObj = ownerToObjMap[dtls['owner']]
        else:
            dmnObj = DaemonObjectsInfo (dtls['owner'], ownerDirsInfo[dtls['owner']], ownerInternalServiceInfo[dtls['owner']]) 
            ownerToObjMap[dtls['owner']] = dmnObj

        dmnObj.objectDict[name] = dtls
    
    for dmn, entry in ownerToObjMap.iteritems():
        entry.parseSrcFile()
        entry.generateThriftInterfaces(ownerToObjMap[dmn])
        entry.createConvertObjToThriftObj(ownerToObjMap[dmn])
        entry.generate_clientif(ownerToObjMap[dmn]) 
        entry.addGeneratedFilesNamesToListing ()
    return


def generateObjectMap():
    genObjInfoJson = JSON_MODEL_REGISTRAION_PATH + 'genObjectConfig.json'
    fd = open(OBJMAP_CODE_GENERATION_PATH + OBJECT_MAP_NAME, 'w+')
    fd.write("""package %s\n\n""" % MODEL_NAME)
    fd.write("""var GenConfigObjectMap = map[string] ConfigObj{\n""")

    with open(genObjInfoJson) as infoFile:
        objData = json.load(infoFile)

    for name,  dtls in objData.iteritems():
        if "w" in dtls['access'] or "r" in dtls['access']:
            line  = "\"%s\" :    &%s{}," %(name, name)
            fd.write(line+"\n")

    fd.write("""}\n""")
    fd.close()
    with open(GENERATED_FILES_LIST, 'a+') as fp:
        fp.write(OBJMAP_CODE_GENERATION_PATH + OBJECT_MAP_NAME+ '\n')

    

if __name__ == "__main__":
    generateObjectMap ()
    generateThriftAndClientIfs()
