#!/usr/bin/env python
import os
import sys
import time
import subprocess

PYANG = 'external/src/pyang/bin/pyang'
PLUGIN_DIR = 'reltools/codegentools/gobind'
GEN_OUT_BASE = 'generated/src/models'
modelsToBuild = [
                 {'srcs'   : '/external/src/openconfig/release/models/interfaces/*.yang',
                  'output' : 'genInterface.go'},

                 {'srcs'   : '/snaproute/src/models/yangmodel/stp/*.yang',
                  'output' : 'stp.go'},

                 {'srcs'   : '/snaproute/src/models/yangmodel/vxlan/*.yang',
                  'output' : 'vxlan.go'},
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
    os.system('mkdir -p ' + srBase + '/' + GEN_OUT_BASE + '/db/')
    for src in modelsToBuild:
        cmd  = srBase + '/' + PYANG + ' --plugindir ' + srBase + '/' + 'reltools/codegentools/gobind/ -f pybind -o ' + \
               srBase + '/' + GEN_OUT_BASE + '/' + src['output'] + ' '  +\
               srBase + src['srcs']

        pyangDir =  srBase + "/external/src/pyang/"
        openConfigDir =  srBase + "/external/src/openconfig/release"
        os.putenv("PYTHONPATH", os.getenv("PYTHONPATH","")+":"+pyangDir)
        yangpath = pyangDir +\
                                 "/modules" + ":" + \
                                 openConfigDir + "/models" + ":" + \
                                 srBase + src['srcs'].rstrip('*.yang') + ":" +\
                                 os.getenv("YANG_MODPATH","")
        print yangpath
        os.putenv("YANG_MODPATH",yangpath)

        os.system(cmd)
