.PHONY: all
all: traffic-web

traffic-web: main.go
	GO111MODULE=on go build

install: all
	systemctl stop traffic-web || true
	cp traffic-web /usr/bin/
	cp traffic-web.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl start traffic-web
	systemctl enable traffic-web
