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

install:installdir
	install $(SRC_DIR)/flexswitch-0.0.1/confd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SRC_DIR)/flexswitch-0.0.1/ribd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SR_CODE_BASE)/snaproute/src/vendors/bcm/opennsl-2.0.0.8-odp/output/cel-x86-trident2-ub14/build/libopennsl.so.1 $(DESTDIR)/opt/$(PROD_NAME)
	install $(SRC_DIR)/flexswitch-0.0.1/asicd $(DESTDIR)/opt/$(PROD_NAME)/
installdir:
	$(MKDIR) $(DESTDIR)
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/bin
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/params


exe: $(COMPDIRS)
	 $(foreach f,$^, make -C $(f) exe DESTDIR=$(DESTDIR);)

ipc: $(COMPS_WITH_IPC)
	 $(foreach f,$^, make -C $(f) ipc DESTDIR=$(DESTDIR);)
