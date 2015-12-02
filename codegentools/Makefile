MKDIR=mkdir -p
DESTDIR=$(SR_CODE_BASE)/snaproute/src/bin
all: codegen installdir

installdir:
	$(MKDIR) $(DESTDIR)

all: codegen
codegen: structs dbif

structs:
	 python $(SR_CODE_BASE)/reltools/codegentools/generateCode.py
dbif:
	 python $(SR_CODE_BASE)/reltools/codegentools/gotodb/sqllite/gotodb.py
