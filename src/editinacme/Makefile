build:
	go build .

install: build
	cp -f ./editinacme /usr/local/bin/
	rm -f ./editinacme

clean:
	rm -f ./editinacme

.PHONY: build install clean
