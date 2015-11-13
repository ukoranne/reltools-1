MKDIR=mkdir -p
DESTDIR=$(SR_CODE_BASE)/snaproute/src/bin
COMPS=asicd\
		config\
		l3

COMPS_WITH_IPC=asicd\
					l3\

all:installdir ipc exe 

installdir:
	$(MKDIR) $(DESTDIR)


exe: $(COMPS)
	 $(foreach f,$^, make -C $(SR_CODE_BASE)/snaproute/src/$(f) exe;)

ipc: $(COMPS_WITH_IPC)
	 $(foreach f,$^, make -C $(SR_CODE_BASE)/snaproute/src/$(f) ipc;)
