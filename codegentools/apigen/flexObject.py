import json
class FlexObject(object) :
    TAB = "    " 
    def __init__ (self,     # Yours truly
                  name,     # Object Name 
                  access,   # Access r/w
                  multiplicity, # UML notation *=many 1=1
                  attrFile): # Location of the attributes description
        self.name = str(name)
        self.access = access
        self.attrFile = attrFile
        self.multiplicity = multiplicity
        self.attrDict = {}
        if name != 'OspfIfEntryConfig':
            return
        with open(attrFile) as fileHdl:
            attrDict = json.load(fileHdl)
            for attrName, tmpInfo in attrDict.iteritems():
                tmpDict = {}
                for k,v in tmpInfo.iteritems():
                    tmpDict[str(k)] = str(v)
                self.attrDict[str(attrName)] = tmpDict
                
                


    def createGetByIdMethod (self, fileHdl):
        print 'Generating GetById Method for %s' %(self.name)

    def createGetMethod (self, fileHdl):
        print 'Generating Get Method for %s' %(self.name)

    def createGetAllMethod (self, fileHdl):
        print 'Generating GetAll Method for %s' %(self.name)

    def writeAllMethods (self, fileHdl):
        self.createGetMethod()

