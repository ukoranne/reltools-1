import json
import re
from flexObject import FlexObject
class FlexConfigObject(FlexObject) :

    def createCreateMethod(self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "def create" + self.name + "(self,"]
        tabs = tabs + self.TAB
        spaces = ' ' * (len(lines[-1])  - len("self, "))
        objLines = [tabs + "obj =  { \n"]
        for attr, attrInfo in self.attrDict.iteritems():
            lines.append("\n" + spaces + "%s," %(attr))
            objLines.append(tabs+tabs + "\'%s\' : %s\n" %(attr, attr))
        lines[-1] = lines[-1][0:lines[-1].find(',')]
        lines.append("):\n")
        objLines.append(tabs + tabs+"}\n")
        lines = lines + objLines
        lines.append (tabs + "reqUrl =  self.urlBase+" +"\'%s\'\n" %(self.name))
        lines.append(tabs + "r = requests.post(reqUrl, data=json.dumps(obj), headers=headers) \n")
        lines.append(tabs + "return r.json() \n\n")
        fileHdl.writelines(lines)

    def createDeleteMethod(self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "def delete" + self.name + "(self,"]
        tabs = tabs + self.TAB
        spaces = ' ' * (len(lines[-1])  - len("self, "))
        objLines = [tabs + "obj =  { \n"]
        for attr, attrInfo in self.attrDict.iteritems():
            lines.append("\n" + spaces + "%s," %(attr))
            objLines.append(tabs+tabs + "\'%s\' : %s\n" %(attr, attr))
        lines[-1] = lines[-1][0:lines[-1].find(',')]
        lines.append("):\n")
        objLines.append(tabs + tabs+"}\n")
        lines = lines + objLines
        lines.append (tabs + "reqUrl =  self.urlBase+" +"\'%s\'\n" %(self.name))
        lines.append(tabs + "r = requests.delete(reqUrl, data=json.dumps(obj), headers=headers) \n")
        lines.append(tabs + "return r.json() \n\n")
        fileHdl.writelines(lines)

    def createDeleteByIdMethod(self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "def delete" + self.name + "ById(self, objectId ):\n"]
        tabs = tabs + self.TAB
        lines.append (tabs + "reqUrl =  self.urlBase+" +"\'%s\'" %(self.name))
        lines[-1] = lines[-1] + "+\"/%s\"%(objectId)\n"
        lines.append(tabs + "r = requests.delete(reqUrl, data=None, headers=headers) \n")
        lines.append(tabs + "return r.json() \n\n")
        fileHdl.writelines(lines)

    def createUpdateMethod (self, fileHdl):
        tabs = self.TAB
        lines = [ "\n"+ tabs + "def update" + self.name + "(self,"]
        tabs = tabs + self.TAB
        spaces = ' ' * (len(lines[-1])  - len("self, "))
        objLines = [tabs + "obj =  { \n"]
        for attr, attrInfo in self.attrDict.iteritems():
            lines.append("\n" + spaces + "%s," %(attr))
            objLines.append(tabs+tabs + "\'%s\' : %s\n" %(attr, attr))
        lines[-1] = lines[-1][0:lines[-1].find(',')]
        lines.append("):\n")
        objLines.append(tabs + tabs+"}\n")
        lines = lines + objLines
        lines.append (tabs + "reqUrl =  self.urlBase+" +"\'%s\'\n" %(self.name))
        lines.append(tabs + "r = requests.patch(reqUrl, data=json.dumps(obj), headers=headers) \n")
        lines.append(tabs + "return r.json() \n\n")
        fileHdl.writelines(lines)

        print 'Generating Update Method for %s' %(self.name)

    def createUpdateByIdMethod (self, fileHdl):
        print 'Generating UpdateById Method for %s' %(self.name)

    def writeAllMethods (self, fileHdl):
        self.createCreateMethod(fileHdl)
        self.createUpdateMethod(fileHdl)
        self.createUpdateByIdMethod(fileHdl)
        self.createDeleteMethod(fileHdl)
        self.createDeleteByIdMethod(fileHdl)
        self.createGetMethod(fileHdl)
        self.createGetByIdMethod(fileHdl)
        self.createGetAllMethod(fileHdl)

