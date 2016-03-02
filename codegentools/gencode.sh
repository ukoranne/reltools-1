#!/bin/bash
cd $SR_CODE_BASE/reltools/codegentools/structs/
source env.sh
pyang --plugindir `pwd` -f pybind  -o $SR_CODE_BASE/snaproute/src/models/ospf.go $SR_CODE_BASE/snaproute/src/models/yangmodel/ospf/OSPF-MIB.yang
pyang --plugindir `pwd` -f pybind  -o $SR_CODE_BASE/snaproute/src/models/stp.go $SR_CODE_BASE/snaproute/src/models/yangmodel/stp/*.yang
mkdir -p  ._genInfo
cd $SR_CODE_BASE/reltools/codegentools/dbif/sqllite
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

