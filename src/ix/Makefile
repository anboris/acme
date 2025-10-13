build:
	go build .

install: build
	go install .
	rm -f ./ix

clean:
	go clean -cache
	rm -f ./ix

.PHONY: build install clean
