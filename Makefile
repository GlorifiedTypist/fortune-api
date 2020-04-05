docs:
	swagger-yaml-to-html.py < pkg/swagger/swagger.yaml > doc/index.html

server:
	rm -rf pkg/swagger/server
	mkdir -p pkg/swagger/server
	go mod init pkg/swagger
	swagger generate server --target pkg/swagger/server --name fortune-api --spec pkg/swagger/swagger.yaml --exclude-main

build:
	go build -o bin/fortune-api internal/main.go internal/warmup.go