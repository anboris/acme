build:
	go build .

install: build
	go install .
	rm -f ./ff

clean:
	go clean -cache
	rm -f ./ff

.PHONY: build install clean
