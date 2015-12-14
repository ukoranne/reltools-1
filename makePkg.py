import os
import json
import subprocess

PACKAGE_BUILD="PKG_BUILD=TRUE"
TEMPLATE_BUILD_TYPE="PKG_BUILD=FALSE"
TEMPLATE_CHANGELOG_VER = "0.0.1"
TEMPLATE_BUILD_DIR = "flexswitch-0.0.1"

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
    preProcess = [
            'cp -a tmplPkgDir ' + build_dir,
            'cp Makefile ' + build_dir,
            'sed -i s/' + TEMPLATE_BUILD_DIR +'/' + build_dir + '/ ' + build_dir +'/Makefile',
            'sed -i s/' + TEMPLATE_BUILD_TYPE +'/' + PACKAGE_BUILD + '/ ' + build_dir + '/Makefile',
            'sed -i s/' + TEMPLATE_CHANGELOG_VER + '/' + parsedPkgInfo['version'] + '/ ' + build_dir + '/debian/changelog',
            ]
    executeCommand(preProcess)
    os.chdir(build_dir)
    pkgRecipe = [
            'fakeroot debian/rules clean',
            'fakeroot debian/rules build',
            'fakeroot debian/rules binary',
            'make clean'
            ]
    executeCommand(pkgRecipe)
    os.chdir("..")
    command = []
    command.append('rm -rf ' + build_dir)
    executeCommand(command)
