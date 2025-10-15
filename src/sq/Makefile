build:
	go build .

install: build
	go install .
	rm -f ./sq

clean:
	go clean -cache
	rm -f ./sq

.PHONY: build install clean
