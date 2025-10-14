CFG := $(filter-out Makefile, $(wildcard *))
DEST := $(HOME)/.config/skhd

.PHONY: all install clean

all: $(CFG)

install: $(CFG)
	@echo installing to $(DEST)
	@for f in $^; do \
		cp $$f $(DEST); \
		echo $$f; \
	done

clean:
	@echo removing from $(DEST)
	@for f in $(CFG); \
		do rm -f $(DEST)/$$f && echo $$f; \
	done
