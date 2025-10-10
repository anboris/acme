CONFIG := config.toml

ifeq ($(shell uname),Darwin)
    DIR := $(HOME)/Library/Application Support/acme-lsp
else
    DIR := $(HOME)/.config/acme-lsp
endif
FILE := $(DIR)/$(CONFIG)

install: $(FILE)

$(FILE): $(CONFIG)
	mkdir -p '$(DIR)'
	cp '$(CONFIG)' '$(FILE)'

.PHONY: install
