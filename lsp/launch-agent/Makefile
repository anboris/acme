NAME := acme-lsp
BIN := $(shell command -v $(NAME) 2>/dev/null || echo $(HOME)/go/bin/$(NAME))
PLIST := 9lab.org.acme-lsp.plist
AGENTD := $(HOME)/Library/LaunchAgents
LOG := /usr/local/var/log/9lab

ifeq ($(shell uname),Darwin)
all: load
else
all:
	@echo "launchd rules are for macOS only"; exit 1
endif

$(LOGD):
	sudo mkdir -p $@
	sudo chown $(USER) $@

$(AGENTD)/$(PLIST): $(PLIST) | $(LOGD)
	mkdir -p $(AGENTD)
	cp $< $@

load: $(AGENTD)/$(PLIST)
	launchctl bootstrap gui/$(shell id -u) $<

unload:
	launchctl unload $(AGENTD)/$(PLIST) 2>/dev/null || true

clean: unload
	rm -f $(AGENTD)/$(PLIST)

.PHONY: all load unload clean
