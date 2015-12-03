MKDIR=mkdir -p
#DESTDIR=$(SR_CODE_BASE)/snaproute/src/bin
#DESTDIR=/usr/bin/snaproute
DESTDIR=$(SR_CODE_BASE)/snaproute/src/flexswitch-0.0.1
SRC_DIR=$(SR_CODE_BASE)/snaproute/src
PROD_NAME=flexswitch

COMPDIRS=$(SRC_DIR)/asicd\
		$(SRC_DIR)/config\
		$(SRC_DIR)/l3

COMPS_WITH_IPC=$(SRC_DIR)/asicd\
					$(SRC_DIR)/l3\

all:installdir ipc exe 

install:installdir copy
	install $(SR_CODE_BASE)/reltools/flexswitch/flexswitch $(DESTDIR)/opt/$(PROD_NAME)
	install $(SR_CODE_BASE)/reltools/flexswitch/daemon.py $(DESTDIR)/opt/$(PROD_NAME)
	install $(SRC_DIR)/flexswitch-0.0.1/confd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SRC_DIR)/flexswitch-0.0.1/arpd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SRC_DIR)/flexswitch-0.0.1/bgpd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SRC_DIR)/flexswitch-0.0.1/ribd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SRC_DIR)/flexswitch-0.0.1/asicd $(DESTDIR)/opt/$(PROD_NAME)/bin

installdir:
	$(MKDIR) $(DESTDIR)
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/kmod
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/bin
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/params
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/sharedlib

copy: $(COMPDIRS)
	 $(foreach f,$^, make -C $(f) install DESTDIR=$(DESTDIR)/opt/$(PROD_NAME);)

exe: $(COMPDIRS)
	 $(foreach f,$^, make -C $(f) exe DESTDIR=$(DESTDIR);)

ipc: $(COMPS_WITH_IPC)
	 $(foreach f,$^, make -C $(f) ipc DESTDIR=$(DESTDIR);)
