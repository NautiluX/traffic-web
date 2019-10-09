.PHONY: all
all: build

build:
	GO111MODULE=on go build

install: build
	cp traffic-web /usr/bin/
	cp traffic-web.service /etc/systemd/system/
	systemctl start traffic-web
	systemctl enable traffic-web
