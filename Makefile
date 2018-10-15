.POSIX:
.SUFFIXES:
.PHONY: clean release debug 

SOURCES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

BINARY=gavage
FULL=github.com/stephane-martin/gavage
VERSION=0.1.0
LDFLAGS=-ldflags '-X main.Version=${VERSION}'
LDFLAGS_RELEASE=-ldflags '-w -s -X main.Version=${VERSION}'

release: ${BINARY}
debug: ${BINARY}_debug

${BINARY}: ${SOURCES}
	go build -a -installsuffix nocgo -tags netgo -o ${BINARY} ${LDFLAGS_RELEASE} ${FULL}

${BINARY}_debug: ${SOURCES}
	go build -x -tags netgo -o ${BINARY}_debug ${LDFLAGS} ${FULL}

clean:
	rm -f ${BINARY} ${BINARY}_debug
	
