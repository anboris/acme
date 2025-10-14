build:
	go build .

install: build
	go install .
	rm -f ./refs

clean:
	go clean -cache
	rm -f ./refs

.PHONY: build install clean
