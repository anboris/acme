BIN := $(filter-out Makefile README, $(wildcard *))
DEST := /usr/local/bin

.PHONY: all install clean

all: $(BIN)

install: $(BIN)
	@echo installing to $(DEST)
	@for f in $^; do \
		cp $$f $(DEST); \
		echo $$f; \
	done

clean:
	@echo removing from $(DEST)
	@for f in $(BIN); \
		do rm -f $(DEST)/$$f && echo $$f; \
	done
