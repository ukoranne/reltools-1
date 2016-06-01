#!/bin/bash
cd $SR_CODE_BASE/reltools/codegentools/structs/
source env.sh
pyang --plugindir `pwd` -f pybind  -o $SR_CODE_BASE/snaproute/src/models/gen_ospf.go $SR_CODE_BASE/snaproute/src/models/yangmodel/ospf/OSPF-MIB.yang
pyang --plugindir `pwd` -f pybind  -o $SR_CODE_BASE/snaproute/src/models/gen_stp.go $SR_CODE_BASE/snaproute/src/models/yangmodel/stp/*.yang
pyang --plugindir `pwd` -f pybind  -o $SR_CODE_BASE/snaproute/src/models/gen_vxlan.go $SR_CODE_BASE/snaproute/src/models/yangmodel/vxlan/*.yang
pyang --plugindir `pwd` -f pybind  -o $SR_CODE_BASE/snaproute/src/models/gen_lacp.go $SR_CODE_BASE/snaproute/src/models/yangmodel/lacp/openconfig-if-aggregate.yang
pyang --plugindir `pwd` -f pybind  -o $SR_CODE_BASE/snaproute/src/models/gen_ldp.go $SR_CODE_BASE/snaproute/src/models/yangmodel/ldp/ldp.yang
mkdir -p  ._genInfo
cd $SR_CODE_BASE/reltools/codegentools/dbif/
./dbifGen.sh
cd $SR_CODE_BASE/reltools/codegentools/thrift
python thriftgen.py
for srcFile in `cat $SR_CODE_BASE/reltools/codegentools/._genInfo/generatedGoFiles.txt`;
do
if [[ $srcFile == *."go"* ]]
then
	   go fmt $srcFile
fi
done    
cd $SR_CODE_BASE/reltools/codegentools/apigen
python generateApis.py 

