MKDIR=mkdir -p
DESTDIR=$(SR_CODE_BASE)/snaproute/src/bin
COMPS=$(SR_CODE_BASE)/snaproute/src/asicd\
		$(SR_CODE_BASE)/snaproute/src/config\
		$(SR_CODE_BASE)/snaproute/src/l3\
		$(SR_CODE_BASE)/snaproute/src/l2/lacp\
		$(SR_CODE_BASE)/snaproute/src/infra/portd

COMPS_WITH_IPC=$(SR_CODE_BASE)/snaproute/src/asicd\
		$(SR_CODE_BASE)/snaproute/src/l3\
		$(SR_CODE_BASE)/snaproute/src/l2/lacp\
		$(SR_CODE_BASE)/snaproute/src/infra/portd

all: codegen installdir ipc exe 

installdir:
	$(MKDIR) $(DESTDIR)


codegen:
	$(MAKE) -f $(SR_CODE_BASE)/reltools/codegentools/Makefile

exe: $(COMPS)
	 $(foreach f,$^, make -C $(f) exe;)

ipc: $(COMPS_WITH_IPC)
	 $(foreach f,$^, make -C $(f) ipc;)

install: $(COMPS)
	 $(foreach f,$^, make -C $(f) ipc;)

clean: $(COMPS)
	 $(foreach f,$^, make -C $(f) clean;)

