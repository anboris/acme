LIB := $(HOME)/lib/plumbing
PLUMBER := plumbing

all: $(LIB)
	cat $(LIB) | 9p write plumb/rules

$(LIB): $(PLUMBER)
	mkdir -p $(dir $(LIB))
	cp $(PLUMBER) $(LIB)
