import json
import re
from flexObject import FlexObject, isNumericAttr, isBoolean, boolFromString
class FlexConfigObject(FlexObject) :

    def createCreateMethod(self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "@processReturnCode"]
        lines.append( "\n"+ tabs + "def create" + self.name + "(self,")
        tabs = tabs + self.TAB
        spaces = ' ' * (len(lines[-1])  - len("self, "))
        objLines = [tabs + "obj =  { \n"]
        for (attr, attrInfo) in self.attrList:
            assignmentStr = ''
            if attrInfo['default'] !="":
                if isNumericAttr(attrInfo['type']):
                    lines.append("\n" + spaces + "%s=%d," %(attr,int(attrInfo['default'].lstrip())))
                    assignmentStr = "int(%s)" %(attr)
                elif isBoolean(attrInfo['type']):
                    lines.append("\n" + spaces + "%s=%s," %(attr, boolFromString(attrInfo['default'].lstrip())))
                    assignmentStr = "True if %s else False" %(attr)
                else:
                    assignmentStr = "%s" %(attr)
                    lines.append("\n" + spaces + "%s=\'%s\'," %(attr,attrInfo['default'].lstrip()))
            else:
                if isNumericAttr(attrInfo['type']):
                    assignmentStr = "int(%s)" %(attr)
                elif isBoolean(attrInfo['type']):
                    assignmentStr = "True if %s else False" %(attr)
                else:
                    assignmentStr = "%s" %(attr)
                lines.append("\n" + spaces + "%s," %(attr))
            objLines.append(tabs+tabs + "\'%s\' : %s,\n" %(attr, assignmentStr))
        lines[-1] = lines[-1][0:lines[-1].find(',')]
        lines.append("):\n")
        objLines.append(tabs + tabs+"}\n")
        lines = lines + objLines
        lines.append (tabs + "reqUrl =  self.urlBase+" +"\'%s\'\n" %(self.name))
        lines.append(tabs + "r = requests.post(reqUrl, data=json.dumps(obj), headers=headers) \n")
        lines.append(tabs + "return r\n")
        fileHdl.writelines(lines)

    def createDeleteMethod(self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "@processReturnCode"]
        lines.append("\n"+ tabs + "def delete" + self.name + "(self,")
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
        lines.append (tabs + "reqUrl =  self.urlBase+" +"\'%s\'\n" %(self.name))
        lines.append(tabs + "r = requests.delete(reqUrl, data=json.dumps(obj), headers=headers) \n")
        lines.append(tabs + "return r\n")
        fileHdl.writelines(lines)

    def createDeleteByIdMethod(self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "@processReturnCode"]
        lines.append("\n"+ tabs + "def delete" + self.name + "ById(self, objectId ):\n")
        tabs = tabs + self.TAB
        lines.append (tabs + "reqUrl =  self.urlBase+" +"\'%s\'" %(self.name))
        lines[-1] = lines[-1] + "+\"/%s\"%(objectId)\n"
        lines.append(tabs + "r = requests.delete(reqUrl, data=None, headers=headers) \n")
        lines.append(tabs + "return r\n")
        fileHdl.writelines(lines)

    def createUpdateMethod (self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "@processReturnCode"]
        lines.append("\n"+ tabs + "def update" + self.name + "(self,")
        tabs = tabs + self.TAB
        spaces = ' ' * (len(lines[-1])  - len("self, "))
        objLines = [tabs + "obj =  {}\n"]
        for (attr, attrInfo) in self.attrList:
            if attrInfo['isKey'] != 'True':
                lines.append("\n" + spaces + "%s = None," %(attr))
            else:
                lines.append("\n" + spaces + "%s," %(attr))
            objLines.append(tabs + "if %s != None :\n" %(attr))
            assignmentStr =''
            if isNumericAttr(attrInfo['type']):
                assignmentStr = "int(%s)" %(attr)
            elif isBoolean(attrInfo['type']):
                assignmentStr = "True if %s else False" %(attr)
            else:
                assignmentStr = "%s" %(attr)
            objLines.append(tabs + self.TAB + "obj[\'%s\'] = %s\n\n" %(attr, assignmentStr))
        lines[-1] = lines[-1][0:lines[-1].find(',')]
        lines.append("):\n")
        lines = lines + objLines
        lines.append (tabs + "reqUrl =  self.urlBase+" +"\'%s\'\n" %(self.name))
        lines.append(tabs + "r = requests.patch(reqUrl, data=json.dumps(obj), headers=headers) \n")
        lines.append(tabs + "return r\n")
        fileHdl.writelines(lines)


    def createUpdateByIdMethod (self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "@processReturnCode"]
        lines.append("\n"+ tabs + "def update" + self.name + "ById(self,\n")
        tabs = tabs + self.TAB
        spaces = ' ' * (len(lines[-1])  - len("self, "))
        lines.append(spaces+ "objectId,")
        objLines = [tabs + "obj =  {\'objectId\': objectId }\n"]
        for (attr, attrInfo) in self.attrList:
            if attrInfo['isKey'] != 'True':
                lines.append("\n" + spaces + "%s = None," %(attr))
                objLines.append(tabs + "if %s !=  None:\n" %(attr))
                objLines.append(tabs + self.TAB + "obj[\'%s\'] = %s\n\n" %(attr, attr))
        lines[-1] = lines[-1][0:lines[-1].find(',')]
        lines.append("):\n")
        lines = lines + objLines
        lines.append (tabs + "reqUrl =  self.urlBase+" +"\'%s\'\n" %(self.name))
        lines.append(tabs + "r = requests.patch(reqUrl, data=json.dumps(obj), headers=headers) \n")
        lines.append(tabs + "return r\n")                                                                                  
        fileHdl.writelines(lines)

    def writeAllMethods (self, fileHdl):
        self.createCreateMethod(fileHdl)
        self.createUpdateMethod(fileHdl)
        self.createUpdateByIdMethod(fileHdl)
        self.createDeleteMethod(fileHdl)
        self.createDeleteByIdMethod(fileHdl)
        self.createGetMethod(fileHdl)
        self.createGetByIdMethod(fileHdl)
        self.createGetAllMethod(fileHdl)

