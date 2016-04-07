import json
import operator
def isNumericAttr (attrInfo) :
    if attrInfo['isArray'] == 'False':
        return attrInfo['type'] in ["int32", "uint32", "uint8"]
    else:
        return False

def isBoolean(attrType) :
    return attrType in ["bool"]

def boolFromString (val) :
    if val == 'false':
        return False
    else:
        return True 

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
        self.attrList = None
        
        with open(attrFile) as fileHdl:
            attrDict = json.load(fileHdl)
            attrList = [None] *len(attrDict)
            keysList = []
            dfltAttrList = []
            for attrName, tmpInfo in attrDict.iteritems():
                tmpDict = {}
                for k,v in tmpInfo.iteritems():
                    tmpDict[str(k)] = str(v)

                self.attrDict[str(attrName)] = tmpDict
                if tmpDict['isKey'] == 'True':
                    keysList.append((attrName, self.attrDict[str(attrName)]))
                elif tmpDict['default'] != '':
                    dfltAttrList.append((attrName, self.attrDict[str(attrName)]))
                else:
                    attrList[tmpInfo['position'] -1] = (attrName, self.attrDict[str(attrName)])

            self.attrList =  keysList + [x for x in attrList if x!= None] + dfltAttrList 

                
    def createGetByIdMethod (self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "@processReturnCode"]
        lines.append("\n"+ tabs + "def get" + self.name + "ById(self, objectId ):\n")
        tabs = tabs + self.TAB
        lines.append (tabs + "reqUrl =  self.stateUrlBase+" +"\'%s\'" %(self.name))
        lines[-1] = lines[-1] + "+\"/%s\"%(objectId)\n"
        lines.append(tabs + "r = requests.get(reqUrl, data=None, headers=headers) \n")
        lines.append(tabs + "return r\n")                                                                                  
        fileHdl.writelines(lines)

    def createGetMethod (self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "@processReturnCode"]
        lines.append("\n"+ tabs + "def get" + self.name + "(self,")
        tabs = tabs + self.TAB
        spaces = ' ' * (len(lines[-1])  - len("self, "))
        objLines = [tabs + "obj =  { \n"]
        for (attr, attrInfo) in self.attrList:
            if attrInfo['isKey'] == 'True':
                lines.append("\n" + spaces + "%s," %(attr))
                objLines.append(tabs+tabs + "\'%s\' : %s,\n" %(attr, attr))
        lines[-1] = lines[-1][0:lines[-1].find(',')]
        lines.append("):\n")
        objLines.append(tabs + tabs+"}\n")
        lines = lines + objLines
        lines.append (tabs + "reqUrl =  self.stateUrlBase+" +"\'%s\'\n" %(self.name))
        lines.append(tabs + "r = requests.get(reqUrl, data=json.dumps(obj), headers=headers) \n")
        lines.append(tabs + "return r\n")                                                                                  
        fileHdl.writelines(lines)

    def createGetAllMethod (self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "def getAll" + self.name+"s" + "(self):\n"]
        tabs = tabs + self.TAB
        lines.append (tabs + "return self.getObjects( \'%s\') \n\n" %(self.name))
        fileHdl.writelines(lines)

    def writeAllMethods (self, fileHdl):
        self.createGetMethod(fileHdl)
        self.createGetByIdMethod(fileHdl)
        self.createGetAllMethod(fileHdl)

