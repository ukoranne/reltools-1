import os
import json
import subprocess

TEMPLATE_BUILD_DIR = "flexswitch-0.0.1"
TEMPLATE_CHANGELOG_VER = "0.0.1"

def executeCommand (command) :
    out = ''
    if type(command) != list:
        command = [ command]
    for cmd in command:
        process = subprocess.Popen(cmd.split(), stdout=subprocess.PIPE)
        out,err = process.communicate()
    return out

if __name__ == '__main__':
    with open("pkgInfo.json", "r") as cfgFile:
        pkgInfo = cfgFile.read().replace('\n', '')
        parsedPkgInfo = json.loads(pkgInfo)
    cfgFile.close()
    build_dir = "flexswitch-" + parsedPkgInfo['version']
    #Create workspace
    command = []
    command.append('cp -a tmplPkgDir ' + build_dir)
    executeCommand(command)
    command = []
    command.append('cp Makefile ' + build_dir)
    executeCommand(command)
    #Edit makefile
    command = []
    command.append('sed -i s/' + TEMPLATE_BUILD_DIR +'/' + build_dir + '/ ' + build_dir +'/Makefile')
    executeCommand(command)
    #Edit changelog
    command = []
    command.append('sed -i s/' + TEMPLATE_CHANGELOG_VER + '/' + parsedPkgInfo['version'] + '/ ' + build_dir + '/debian/changelog')
    executeCommand(command)
    #Build package
    command = []
    os.chdir(build_dir)
    command.append('fakeroot debian/rules clean')
    executeCommand(command)
    command = []
    command.append('fakeroot debian/rules build')
    executeCommand(command)
    command = []
    command.append('fakeroot debian/rules binary')
    executeCommand(command)
    #Cleanup build dir
    os.chdir("..")
    command = []
    command.append('rm -rf ' + build_dir)
    executeCommand(command)
