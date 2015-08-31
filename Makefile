EXECUTABLE := mysql_probe

GO_SRC := \
  main.go \
	mysqltest/**/*.go

UNIX_EXECUTABLES := \
	darwin/386/$(EXECUTABLE) \
	darwin/amd64/$(EXECUTABLE) \
	linux/386/$(EXECUTABLE) \
	linux/amd64/$(EXECUTABLE)

COMPRESSED_EXECUTABLES=$(UNIX_EXECUTABLES:%=%.tar.bz2)
COMPRESSED_EXECUTABLE_TARGETS=$(COMPRESSED_EXECUTABLES:%=bin/%)

build: $(EXECUTABLE)

# arm
bin/linux/arm/5/$(EXECUTABLE): $(GO_SRC)
	GOARM=5 GOARCH=arm GOOS=linux go build -o "$@"
bin/linux/arm/7/$(EXECUTABLE): $(GO_SRC)
	GOARM=7 GOARCH=arm GOOS=linux go build -o "$@"

# 386
bin/darwin/386/$(EXECUTABLE): $(GO_SRC)
	GOARCH=386 GOOS=darwin go build -o "$@"
bin/linux/386/$(EXECUTABLE): $(GO_SRC)
	GOARCH=386 GOOS=linux go build -o "$@"
bin/windows/386/$(EXECUTABLE): $(GO_SRC)
	GOARCH=386 GOOS=windows go build -o "$@"

# amd64
bin/freebsd/amd64/$(EXECUTABLE): $(GO_SRC)
	GOARCH=amd64 GOOS=freebsd go build -o "$@"
bin/darwin/amd64/$(EXECUTABLE): $(GO_SRC)
	GOARCH=amd64 GOOS=darwin go build -o "$@"
bin/linux/amd64/$(EXECUTABLE): $(GO_SRC)
	GOARCH=amd64 GOOS=linux go build -o "$@"
bin/windows/amd64/$(EXECUTABLE).exe: $(GO_SRC)
	GOARCH=amd64 GOOS=windows go build -o "$@"

# compressed artifacts, makes a huge difference (Go executable is ~9MB,
# after compressing ~2MB)
%.tar.bz2: %
	tar -jcvf "$<.tar.bz2" "$<"
%.zip: %.exe
	zip "$@" "$<"

# install and/or update all dependencies, run this from the project directory
# go get -u ./...
# go test -i ./
dep:
	go list -f '{{join .Deps "\n"}}' | xargs go list -e -f '{{if not .Standard}}{{.ImportPath}}{{end}}' | grep -v 'github.com/haikulearning/mysql_probe' | xargs go get -u

$(EXECUTABLE): dep $(GO_SRC)
	go build -o "$@"

install:
	go install

all: $(EXECUTABLE) osx linux

osx: bin/darwin/386/$(EXECUTABLE) bin/darwin/amd64/$(EXECUTABLE)

linux: bin/linux/386/$(EXECUTABLE) bin/linux/amd64/$(EXECUTABLE)

release: $(COMPRESSED_EXECUTABLE_TARGETS)
	git push && git push --tags

clean:
	rm $(EXECUTABLE) || true
	rm -rf bin/

.PHONY: clean dep install all osx linux