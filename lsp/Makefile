CONFIG := config.toml
BINS := $(wildcard bin/*)

ifeq ($(shell uname),Darwin)
    DIR := $(HOME)/Library/Application Support/acme-lsp
    BIN_DIR := /usr/local/bin
else
    DIR := $(HOME)/.config/acme-lsp
    BIN_DIR := $(HOME)/.local/bin
endif

FILE := $(DIR)/$(CONFIG)

.PHONY: install clean

install: $(FILE)
	mkdir -p '$(BIN_DIR)'
	@for bin in $(BINS); do \
		echo cp "$$bin" '$(BIN_DIR)'; \
		install -m 755 "$$bin" '$(BIN_DIR)'; \
	done

$(FILE): $(CONFIG)
	mkdir -p '$(DIR)'
	cp '$(CONFIG)' '$(FILE)'

BIN_NAMES = $(notdir $(BINS))

clean:
	rm -f '$(FILE)'
	@for name in $(BIN_NAMES); do \
		echo rm -f '$(BIN_DIR)/'$$name; \
		rm -f '$(BIN_DIR)/'$$name; \
	done
