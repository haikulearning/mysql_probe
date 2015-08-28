mysql_probe: main.go mysqltest/**/*.go
	go build

all: mysql_probe osx linux

osx: main.go mysqltest/**/*.go
	GOOS=darwin GOARCH=386 go build -o mysql_probe.mac.386
	GOOS=darwin GOARCH=amd64 go build -o mysql_probe.mac.amd64

linux: main.go mysqltest/**/*.go
	GOOS=linux GOARCH=386 go build -o mysql_probe.linux.386
	GOOS=linux GOARCH=amd64 go build -o mysql_probe.linux.amd64

clean:
	rm -f mysql_probe mysql_probe.*
