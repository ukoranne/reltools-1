MKDIR=mkdir -p
RMDIRFORCE=rm -rf
PROD_NAME=flexswitch
BUILD_DIR=flexswitch-0.0.1
SRCDIR=$(SR_CODE_BASE)/snaproute/src
DESTDIR=$(SR_CODE_BASE)/snaproute/src/$(BUILD_DIR)
COMPS=$(SR_CODE_BASE)/snaproute/src/asicd\
		$(SR_CODE_BASE)/snaproute/src/config\
		$(SR_CODE_BASE)/snaproute/src/infra\
		$(SR_CODE_BASE)/snaproute/src/l3

COMPS_WITH_IPC=$(SRCDIR)/asicd\
				$(SRCDIR)/l3\
				$(SRCDIR)/infra

#FIXME: Add codegen once things are stable
#all: codegen installdir ipc exe install
all: installdir ipc exe install

installdir:
	$(MKDIR) $(DESTDIR)
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/kmod
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/bin
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/params
	$(MKDIR) $(DESTDIR)/opt/$(PROD_NAME)/sharedlib

codegen:
	$(MAKE) -f $(SR_CODE_BASE)/reltools/codegentools/Makefile

exe: $(COMPS)
	$(foreach f,$^, make -C $(f) exe DESTDIR=$(DESTDIR);)

ipc: $(COMPS_WITH_IPC)
	$(foreach f,$^, make -C $(f) ipc DESTDIR=$(DESTDIR);)

copy: $(COMPS)
	$(foreach f,$^, make -C $(f) install DESTDIR=$(DESTDIR)/opt/$(PROD_NAME);)

install:installdir copy
	install $(SR_CODE_BASE)/reltools/flexswitch $(DESTDIR)/opt/$(PROD_NAME)
	install $(SR_CODE_BASE)/reltools/daemon.py $(DESTDIR)/opt/$(PROD_NAME)
	install $(SRCDIR)/$(BUILD_DIR)/confd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SRCDIR)/$(BUILD_DIR)/arpd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SRCDIR)/$(BUILD_DIR)/bgpd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SRCDIR)/$(BUILD_DIR)/ribd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SRCDIR)/$(BUILD_DIR)/asicd $(DESTDIR)/opt/$(PROD_NAME)/bin
	install $(SR_CODE_BASE)/external/src/github.com/nanomsg/nanomsg/.libs/libnanomsg.so.4.0.0 $(DESTDIR)/opt/$(PROD_NAME)/sharedlib

clean: $(COMPS)
	$(foreach f,$^, make -C $(f) clean;)
	$(RMDIRFORCE) $(DESTDIR)
