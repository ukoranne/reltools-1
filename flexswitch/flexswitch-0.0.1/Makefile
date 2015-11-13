MKDIR=mkdir -p
#DESTDIR=$(SR_CODE_BASE)/snaproute/src/bin
#DESTDIR=/usr/bin/snaproute
DESTDIR=/home/hari/git/snaproute/src/flexswitch-0.0.1
SRC_DIR=$(SR_CODE_BASE)/snaproute/src

COMPDIRS=$(SRC_DIR)/asicd\
		$(SRC_DIR)/config\
		$(SRC_DIR)/l3

COMPS_WITH_IPC=$(SRC_DIR)/asicd\
					$(SRC_DIR)/l3\

all:installdir ipc exe 

install:installdir
	install $(SRC_DIR)/flexswitch-0.0.1/confd $(DESTDIR)/opt/snaproute/
	install $(SRC_DIR)/flexswitch-0.0.1/ribd $(DESTDIR)/opt/snaproute/
	#install $(SRC_DIR)/flexswitch-0.0.1/asicd $(DESTDIR)

installdir:
	$(MKDIR) $(DESTDIR)
	$(MKDIR) $(DESTDIR)/opt/snaproute/


exe: $(COMPDIRS)
	 $(foreach f,$^, make -C $(f) exe DESTDIR=$(DESTDIR);)

ipc: $(COMPS_WITH_IPC)
	 $(foreach f,$^, make -C $(f) ipc DESTDIR=$(DESTDIR);)
