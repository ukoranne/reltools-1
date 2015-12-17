MKDIR=mkdir -p
RMDIRFORCE=rm -rf
PKG_BUILD=TRUE
PROD_NAME=flexswitch
ifneq (,$(findstring $(PKG_BUILD), FALSE))
	EXT_INSTALL_PATH=
	BUILD_DIR=out
	ALL_DEPS=installdir ipc exe
else
	EXT_INSTALL_PATH=/opt/$(PROD_NAME)
	BUILD_DIR=flexswitch-0.0.5
	ALL_DEPS=installdir ipc exe install
endif
SRCDIR=$(SR_CODE_BASE)/snaproute/src
DESTDIR=$(SR_CODE_BASE)/snaproute/src/$(BUILD_DIR)
ifneq (,$(findstring $(PKG_BUILD), FALSE))
	EXE_DIR=/bin
else
	EXE_DIR=
endif
COMPS=$(SR_CODE_BASE)/snaproute/src/asicd\
		$(SR_CODE_BASE)/snaproute/src/config\
		$(SR_CODE_BASE)/snaproute/src/infra\
		$(SR_CODE_BASE)/snaproute/src/l3\
		$(SR_CODE_BASE)/snaproute/src/l2

COMPS_WITH_IPC=$(SR_CODE_BASE)/snaproute/src/asicd\
		$(SR_CODE_BASE)/snaproute/src/l3\
		$(SR_CODE_BASE)/snaproute/src/l2\
		$(SR_CODE_BASE)/snaproute/src/infra

#FIXME: Add codegen once things are stable
#all: codegen installdir ipc exe install
all: $(ALL_DEPS)

installdir:
	$(MKDIR) $(DESTDIR)
	$(MKDIR) $(DESTDIR)/$(EXT_INSTALL_PATH)/
	$(MKDIR) $(DESTDIR)/$(EXT_INSTALL_PATH)/kmod
	$(MKDIR) $(DESTDIR)/$(EXT_INSTALL_PATH)/bin
	$(MKDIR) $(DESTDIR)/$(EXT_INSTALL_PATH)/params
	$(MKDIR) $(DESTDIR)/$(EXT_INSTALL_PATH)/sharedlib

codegen:
	$(MAKE) -f $(SR_CODE_BASE)/reltools/codegentools/Makefile

exe: $(COMPS)
	$(foreach f,$^, make -C $(f) exe DESTDIR=$(DESTDIR)/$(EXE_DIR) GOLDFLAGS="-r /opt/flexswitch/sharedlib";)

ipc: $(COMPS_WITH_IPC)
	$(foreach f,$^, make -C $(f) ipc DESTDIR=$(DESTDIR);)

copy: $(COMPS)
	$(foreach f,$^, make -C $(f) install DESTDIR=$(DESTDIR)/$(EXT_INSTALL_PATH);)

install:installdir copy
	install $(SR_CODE_BASE)/reltools/flexswitch $(DESTDIR)/$(EXT_INSTALL_PATH)
	install $(SR_CODE_BASE)/reltools/daemon.py $(DESTDIR)/$(EXT_INSTALL_PATH)
	install $(SRCDIR)/$(BUILD_DIR)/confd $(DESTDIR)/$(EXT_INSTALL_PATH)/bin
	install $(SRCDIR)/$(BUILD_DIR)/arpd $(DESTDIR)/$(EXT_INSTALL_PATH)/bin
	install $(SRCDIR)/$(BUILD_DIR)/bgpd $(DESTDIR)/$(EXT_INSTALL_PATH)/bin
	install $(SRCDIR)/$(BUILD_DIR)/ribd $(DESTDIR)/$(EXT_INSTALL_PATH)/bin
	install $(SRCDIR)/$(BUILD_DIR)/portd $(DESTDIR)/$(EXT_INSTALL_PATH)/bin
	install $(SRCDIR)/$(BUILD_DIR)/asicd $(DESTDIR)/$(EXT_INSTALL_PATH)/bin
	install $(SR_CODE_BASE)/external/src/github.com/nanomsg/nanomsg/.libs/libnanomsg.so.4.0.0 $(DESTDIR)/$(EXT_INSTALL_PATH)/sharedlib

clean: $(COMPS)
	$(foreach f,$^, make -C $(f) clean;)
	$(RMDIRFORCE) $(DESTDIR)
