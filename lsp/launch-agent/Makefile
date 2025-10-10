NAME := acme-lsp
BIN := $(shell command -v $(NAME) 2>/dev/null || echo $(HOME)/go/bin/$(NAME))
AGENTD := $(HOME)/Library/LaunchAgents
LOG := /usr/local/var/log/9lab

PLIST1 := 9lab.org.acme-lsp.plist
PLIST2 := 9lab.org.acmefocused.plist

ifeq ($(shell uname),Darwin)
all: load
else
all:
	@echo "launchd rules are for macOS only"; exit 1
endif

$(LOG):
	sudo mkdir -p $@
	sudo chown $(USER) $@

$(AGENTD)/$(PLIST1): $(PLIST1) | $(LOG)
	mkdir -p $(AGENTD)
	cp $< $@

$(AGENTD)/$(PLIST2): $(PLIST2) | $(LOG)
	mkdir -p $(AGENTD)
	cp $< $@

load: $(AGENTD)/$(PLIST1) $(AGENTD)/$(PLIST2)
	launchctl bootstrap gui/$(shell id -u) $(AGENTD)/$(PLIST1)
	launchctl bootstrap gui/$(shell id -u) $(AGENTD)/$(PLIST2)

unload:
	launchctl unload $(AGENTD)/$(PLIST1) 2>/dev/null || true
	launchctl unload $(AGENTD)/$(PLIST2) 2>/dev/null || true

clean: unload
	rm -f $(AGENTD)/$(PLIST1) $(AGENTD)/$(PLIST2)

.PHONY: all load unload clean
