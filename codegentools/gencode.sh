#!/bin/bash
cd $SR_CODE_BASE/reltools/codegentools/structs/
source env.sh
pyang --plugindir `pwd` -f pybind  -o $SR_CODE_BASE/snaproute/src/models/ospf.go $SR_CODE_BASE/snaproute/src/models/yangmodel/ospf/OSPF-MIB.yang
cd $SR_CODE_BASE/reltools/codegentools/structs/
cd $SR_CODE_BASE/reltools/codegentools/dbif/sqllite
./dbifGen.sh

