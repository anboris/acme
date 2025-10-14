DEST := $(HOME)/.config/skhd
CFG := skhdrc

.PHONY: all install clean

all: install

install: $(DEST)/$(CFG)

$(DEST)/$(CFG): $(CFG)
	mkdir -p $(DEST)
	cp $< $@

clean:
	rm -f $(DEST)/$(CFG)
