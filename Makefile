COMPS=asicd\
		config\
		l3

COMPS_WITH_IPC=asicd\
					l3\

all:ipc exe 

exe: $(COMPS)
	 $(foreach f,$^, make -C $(f) exe;)

ipc: $(COMPS_WITH_IPC)
	 $(foreach f,$^, make -C $(f) ipc;)
