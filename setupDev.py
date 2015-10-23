import os
import subprocess
thrift_version = '0.9.3'
thrift_pkg_name = 'thrift-'+thrift_version 
thrift_tar = thrift_pkg_name +'.tar.gz'

TMP_DIR = ".tmp"
SNAP_ROUTE_SRC = '/snaproute/src/'
EXTERNAL_SRC = '/external/src/'

gHomeDir = None
gDryRun =  False

def executeCommand (command) :
    out = ''
    if type(command) != list:
        command = [ command]
    for cmd in command:
        if gDryRun :
            print cmd
        else:
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
    #return True
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

def gitRepoSyncToTag(dirloc, tag):
    os.chdir(dirloc)
    command = 'git checkout tags/'+ tag
    executeCommand(command)
    
def setRemoteUpstream (repoUrl):
    command = 'git remote add upstream ' + repoUrl
    executeCommand(command)

def getGolangExternalDependencies(repourl, dirloc):
    os.chdir(dirloc)
    command = 'git clone '+ repourl
    executeCommand(command)

def createDirectoryStructure() :
    dirs = [SNAP_ROUTE_SRC,EXTERNAL_SRC]
    for everydir in dirs:
        command = 'mkdir -p '+ gHomeDir + everydir 
        executeCommand(command)

def getExternalGoDeps() :
    externalGoDeps = [
                     { 'repo'      : 'thrift',
                       'reltag'    : '0.9.3',
                       'renamesrc' : 'thrift',
                       'renamedst' : 'git.apache.org/thrift.git'
                     },
                     { 'repo'       : 'mux',
                       'renamesrc'  : 'mux',
                       'renamedst'  : 'github.com/gorilla/'
                     },
                     { 'repo'       : 'context',
                       'renamesrc'  : 'context',
                       'renamedst'  : 'github.com/gorilla/'
                     }
                     ]

    dirLocation = gHomeDir + EXTERNAL_SRC 
    for dep in externalGoDeps:
        repoUrl = 'https://github.com/SnapRoute/'+ dep['repo']
        cloneGitRepo ( repoUrl , dirLocation)
        if dep.has_key('reltag'):
            gitRepoSyncToTag(dirLocation+dep['repo'], dep['reltag'])
        dstDir = dep['renamedst']
        dirToMake = dstDir 
        if not dstDir.endswith('/'):
            dirToMake = ''.join(dstDir.split('/')[:-1])
        os.chdir(dirLocation)
        for d in dirToMake.split('/'):
            if len(d):
                cmd  =  'mkdir -p ' + d
                executeCommand(cmd)
                os.chdir(d)
        cmd = 'mv ' + dirLocation + dep['renamesrc']+ ' ' + dirLocation + dep['renamedst']
        executeCommand(cmd)

def cloneSnapRouteGitRepos():
    userRepoPrefix   = 'https://github.com/'+gUserName+'/'
    remoteRepoPrefix = 'https://github.com/'+ 'SnapRoute/'
    dirLocation      = gHomeDir + SNAP_ROUTE_SRC

    gitReposToClone = [ 'l2', 'l3', 'utils', 'asicd', 'config', 'models'] # (URL, DIR)
                        #('https://'+gUserName+'@github.com/SnapRoute/thrift', gHomeDir+EXTERNAL_SRC),
    for repo in gitReposToClone:
        cloneGitRepo ( userRepoPrefix + repo , dirLocation)
        setRemoteUpstream (remoteRepoPrefix +repo+'.git')

if __name__ == '__main__':
    gUserName =  raw_input('Please Enter github username:')
    gHomeDir = os.path.dirname(os.getcwd())
    print '### Anchor Directory is %s' %(gHomeDir)

    createDirectoryStructure()
    if False == verifyThriftInstallation():
        installThriftDependencies()
        installThrift()
    else:
        print ' Thrift already exists'

    cloneSnapRouteGitRepos()
    getExternalGoDeps()
