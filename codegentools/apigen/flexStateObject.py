from flexObject import FlexObject
class FlexStateObject(FlexObject) :
    def createGetByIdMethod (self, fileHdl):
        print 'Generating GetById Method for %s' %(self.name)

    def createGetMethod (self, fileHdl):
        print 'Generating Get Method for %s' %(self.name)

    def createGetAllMethod (self, fileHdl):
        print 'Generating GetAll Method for %s' %(self.name)

    def writeAllMethods (self, fileHdl):
        self.createGetMethod(fileHdl)
        self.createGetByIdMethod(fileHdl)
        self.createGetAllMethod(fileHdl)

