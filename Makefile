all: build

build: clean
	CGO_ENABLED=0 go build -o /go/bin/apitest

clean:
	rm -rf /go/bin/apitest

.PHONY: all build clean
