PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

docs:
	swagger-yaml-to-html.py < pkg/swagger/swagger.yaml > doc/index.html

server:
	rm -rf pkg/swagger/server
	mkdir -p pkg/swagger/server
	go mod init pkg/swagger
	swagger generate server --target pkg/swagger/server --name fortune-api --spec pkg/swagger/swagger.yaml --exclude-main

build:
	go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go || exit

production:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go

