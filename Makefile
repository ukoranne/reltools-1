MKDIR=mkdir -p
DESTDIR=$(SR_CODE_BASE)/snaproute/src/bin
COMPS=$(SR_CODE_BASE)/snaproute/src/asicd\
		$(SR_CODE_BASE)/snaproute/src/config\
		$(SR_CODE_BASE)/snaproute/src/infra\
		$(SR_CODE_BASE)/snaproute/src/l3

COMPS_WITH_IPC=$(SR_CODE_BASE)/snaproute/src/asicd\
					$(SR_CODE_BASE)/snaproute/src/l3\
					$(SR_CODE_BASE)/snaproute/src/infra

all: codegen installdir ipc exe install

installdir:
	$(MKDIR) $(DESTDIR)


codegen:
	$(MAKE) -f $(SR_CODE_BASE)/reltools/codegentools/Makefile

exe: $(COMPS)
	 $(foreach f,$^, make -C $(f) exe;)

ipc: $(COMPS_WITH_IPC)
	 $(foreach f,$^, make -C $(f) ipc;)

install: $(COMPS)
	 $(MKDIR) $(DESTDIR)/kmod
	 $(foreach f,$^, make -C $(f) install;)

clean: $(COMPS)
	 $(foreach f,$^, make -C $(f) clean;)

