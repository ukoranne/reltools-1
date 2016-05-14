import os
import json
import subprocess
from optparse import OptionParser

PACKAGE_BUILD="PKG_BUILD=TRUE"
TEMPLATE_BUILD_TYPE="PKG_BUILD=FALSE"
TEMPLATE_CHANGELOG_VER = "0.0.1"
TEMPLATE_BUILD_DIR = "flexswitch-0.0.1"
TEMPLATE_SAI_TARGET = "mlnx"
TEMPLATE_OPENNSL_TARGET = "cel_redstone"

def executeCommand (command) :
    out = ''
    if type(command) != list:
        command = [ command]
    for cmd in command:
        process = subprocess.Popen(cmd.split(), stdout=subprocess.PIPE)
        out,err = process.communicate()
    return out

if __name__ == '__main__':
    #Parse options
    parser = OptionParser()
    parser.add_option("-s", "--sai",
                      dest="saiTarget",
                      default="mlnx",
                      action='store',
                      help="Taget platform to use for SAI plugin (mlnx/bfoot)")
    parser.add_option("-o", "--opennsl",
                      dest="opennslTarget",
                      default="cel_redstone",
                      action='store',
                      help="Target platform to use for opennsl plugin (currently unused)")

    cmd = 'python  buildInfoGen.py'
    executeCommand(cmd)
    with open("pkgInfo.json", "r") as cfgFile:
        pkgInfo = cfgFile.read().replace('\n', '')
        parsedPkgInfo = json.loads(pkgInfo)
    cfgFile.close()
    pkgVersion = parsedPkgInfo['major']+ '.' + parsedPkgInfo['minor'] +  '.' + parsedPkgInfo['patch'] + '.' + parsedPkgInfo['build']
    build_dir = "flexswitch-" + pkgVersion
    preProcess = [
            'cp -a tmplPkgDir ' + build_dir,
            'cp Makefile ' + build_dir,
            'sed -i s/' + TEMPLATE_BUILD_DIR +'/' + build_dir + '/ ' + build_dir +'/Makefile',
            'sed -i s/' + TEMPLATE_BUILD_TYPE +'/' + PACKAGE_BUILD + '/ ' + build_dir + '/Makefile',
            'sed -i s/' + TEMPLATE_CHANGELOG_VER + '/' + pkgVersion+ '/ ' + build_dir + '/debian/changelog',
            ]
    executeCommand(preProcess)

    #Override sai, opennsl targets
    (opts, args) = parser.parse_args()
    if opts.saiTarget != None:
        saiTarget = opts.saiTarget
        executeCommand('sed -i s/' + TEMPLATE_SAI_TARGET +'/' + saiTarget + '/ ' + build_dir + '/Makefile')
    if opts.opennslTarget != None:
        opennslTarget = opts.opennslTarget
        executeCommand('sed -i s/' + TEMPLATE_OPENNSL_TARGET +'/' + opennslTarget + '/ ' + build_dir + '/Makefile')

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
