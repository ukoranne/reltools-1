import os
import subprocess
thrift_version = '0.9.3'
thrift_pkg_name = 'thrift-'+thrift_version 
thrift_tar = thrift_pkg_name +'.tar.gz'

TMP_DIR = ".tmp"
SNAP_ROUTE_SRC = '/snaproute/src/'
EXTERNAL_SRC = '/external/src/github.com/'

gHomeDir = None

def executeCommand (command) :
    if type(command) != list:
        command = [ command]
    for cmd in command:
        #print 'Executing command %s' %(cmd)
        #print "%s" %(cmd.split())
        #print subprocess.check_output(cmd.split())
        process = subprocess.Popen(cmd.split(), stdout=subprocess.PIPE)
        out,err = process.communicate()
        return out

def downloadThrift() :
    print 'Downloading Thrift'
    command = []
    command.append('mkdir -p '+TMP_DIR)
    executeCommand(command)

    command = []
    command.append('wget -O '+ TMP_DIR + '/' +thrift_tar+ ' '+ 
                    'http://apache.arvixe.com/thrift/0.9.3/thrift-0.9.3.tar.gz')
    executeCommand(command)

def verifyThriftInstallation():
    command = []
    command.append('thrift -version')
     #'Thrift version 0.9.3'
    resp = executeCommand(command)
    print 'Thrift version check returned %s' %(resp)
    if thrift_version in resp:
        return True
    else:
        return False

def installThrift() :

    if True == verifyThriftInstallation():
        print ' Thrift already exists returning'
        return

    downloadThrift()
    os.chdir(TMP_DIR)
    command = 'tar -xvf '+ thrift_tar
    executeCommand(command)

    os.chdir(thrift_pkg_name)

    command = './configure'
    executeCommand(command)

    command = 'make'
    executeCommand(command)

    command = 'sudo make install'
    executeCommand(command)

    if False== verifyThriftInstallation():
        print ' Thrift Installation failed'
        return

def installThriftDependencies ():
    command = []
    command.append('sudo apt-get install libboost-dev libboost-test-dev libboost-program-options-dev libboost-system-dev libboost-filesystem-dev libevent-dev automake libtool flex bison pkg-config g++ libssl-dev ant')
    executeCommand(command)

def cloneGitRepo (repourl, dirloc):
    os.chdir(dirloc)
    command = 'git clone '+ repourl
    executeCommand(command)
    
def getGolandExternalDependencies(repourl, dirloc):
    os.chdir(dirloc)
    command = 'git clone '+ repourl
    executeCommand(command)

def createDirectoryStructure() :
    dirs = [SNAP_ROUTE_SRC,EXTERNAL_SRC]
    for everydir in dirs:
        command = 'mkdir -p '+ gHomeDir + everydir 
        executeCommand(command)

if __name__ == '__main__':
    gUserName =  raw_input('Please Enter github username:')
    gHomeDir =  os.getcwd()
    #print '### Home Directory is %s' %(gHomeDir)

    createDirectoryStructure()
    if False== verifyThriftInstallation():
        print ' Thrift doesnt exist'
        installThriftDependencies()
        installThrift()
    else:
        print ' Thrift already exists'
    gitReposToClone = [ # (URL, DIR)
                        ('https://'+gUserName+'@github.com/SnapRoute/thrift', gHomeDir+EXTERNAL_SRC),
                        ('https://'+gUserName+'@github.com/'+gUserName+'/l2', gHomeDir+SNAP_ROUTE_SRC),
                        ('https://'+gUserName+'@github.com/'+gUserName+'/l3', gHomeDir+SNAP_ROUTE_SRC),
                        ('https://'+gUserName+'@github.com/'+gUserName+'/utils', gHomeDir+SNAP_ROUTE_SRC),
                        ('https://'+gUserName+'@github.com/'+gUserName+'/asicd', gHomeDir+SNAP_ROUTE_SRC),
                        ('https://'+gUserName+'@github.com/'+gUserName+'/config', gHomeDir+SNAP_ROUTE_SRC),
                        ('https://'+gUserName+'@github.com/'+gUserName+'/models', gHomeDir+SNAP_ROUTE_SRC),
                      ]
    for repo in gitReposToClone:
        #cloneGitRepo(repo[0], repo[1])
        pass

