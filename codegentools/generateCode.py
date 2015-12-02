#!/usr/bin/env python
import os
import sys
import time
import subprocess

PYANG = 'external/src/pyang/bin/pyang'
PLUGIN_DIR = 'reltools/codegentools/gobind'
GEN_OUT_BASE = 'generated/src/model'
modelsToBuild = [
                 {'srcs'   : 'openconfig/release/models/interfaces/*.yang',
                  'output' : 'genInterface.go'},

                 #{'srcs'   : 'openconfig/release/models/vlan/*.yang',
                 # 'output' : 'genVlan.go'},

                 #{'srcs'   : 'openconfig/release/models/bgp/*.yang',
                 # 'output' : 'genBgp.go'},

                ]

# Command looks like this
#pyang --plugindir ~/git/reltools/codegentools/gobind/ -f pybind -o interface.go *.yang
if __name__=="__main__":
    srBase = os.environ.get('SR_CODE_BASE', None)

    if srBase is None:
        print '****** Error ******'
        print 'Hate to blame you but SR_CODE_BASE is missing.\nDid you forget to set the environment variables'
        print 'Try adding line \"export SR_CODE_BASE=$HOME/git\" to bashrc'
        sys.exit(0)

    os.system('mkdir -p ' + srBase + '/' + GEN_OUT_BASE + '/')
    for src in modelsToBuild:
        cmd  = srBase + '/' + PYANG + ' --plugindir ' + srBase + '/' + 'reltools/codegentools/gobind/ -f pybind -o ' + \
               srBase + '/' + GEN_OUT_BASE + '/' + src['output'] + ' '  +\
               srBase + '/external/src/' + src['srcs']

        pyangDir =  srBase + "/external/src/pyang/"
        openConfigDir =  srBase + "/external/src/openconfig/release"
        os.putenv("PYTHONPATH", os.getenv("PYTHONPATH","")+":"+pyangDir)
        os.putenv("YANG_MODPATH", pyangDir +\
                                 "/modules" + ":" + \
                                 openConfigDir + "/models" +\
                                 os.getenv("YANG_MODPATH","")
                                 )

        os.system(cmd)
