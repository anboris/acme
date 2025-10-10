build:
	go build .

install: build
	go install .
	rm -f ./gg

clean:
	go clean -cache
	rm -f ./gg

.PHONY: build install clean
