NAME := acme-lsp
BIN := $(shell command -v $(NAME) 2>/dev/null || echo $(HOME)/go/bin/$(NAME))
AGENTD := $(HOME)/Library/LaunchAgents
LOG := /usr/local/var/log/9lab

PLISTS := 9lab.org.acme-lsp.plist 9lab.org.acmefocused.plist
PLIST_PATHS := $(addprefix $(AGENTD)/,$(PLISTS))

ifeq ($(shell uname),Darwin)
all: load
else
all:
	@echo "launchd rules are for macOS only"; exit 1
endif

$(LOG):
	@sudo mkdir -p $@
	@sudo chown $(USER) $@
	@echo "Created log directory: $@"

$(AGENTD):
	@mkdir -p $@
	@echo "Ensured LaunchAgents directory: $@"

$(AGENTD)/%: % | $(AGENTD) $(LOG)
	@cp $< $@
	@echo "Installed $< → $@"

load: $(PLIST_PATHS)
	@for plist in $(PLIST_PATHS); do \
		echo "Loading $$plist"; \
		launchctl load -w $$plist || echo "Failed to load $$plist"; \
	done

unload:
	@for plist in $(PLIST_PATHS); do \
		echo "Unloading $$plist"; \
		launchctl unload -w $$plist 2>/dev/null || true; \
	done

clean: unload
	@rm -f $(PLIST_PATHS)
	@echo "Cleaned up LaunchAgents."

.PHONY: all load unload clean
