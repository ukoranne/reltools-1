MKDIR=mkdir -p
DESTDIR=$(SR_CODE_BASE)/snaproute/src/bin
all: codegen installdir

installdir:
	$(MKDIR) $(DESTDIR)

all: codegen
codegen: structs json thriftclient dbif

structsv1:
	python $(SR_CODE_BASE)/reltools/codegentools/generateCode.py
dbifv1:
	python $(SR_CODE_BASE)/reltools/codegentools/gotodb/sqllite/gotodb.py
json:
	python $(SR_CODE_BASE)/reltools/codegentools/gotojson/gotojson.py
thriftclient:
	python $(SR_CODE_BASE)/reltools/codegentools/gotothrift/gotothrift.py

clean:
	rm -rf $(SR_CODE_BASE)/generated/src/models/
	rm -rf $(SR_CODE_BASE)/generated/src/config/
	rm -rf $(SR_CODE_BASE)/generated/src/gorpc/
	
	
