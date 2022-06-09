export tag=v1.0
root:
	export ROOT=github.com/countryhu/cloudnativecamp

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t cloudnativecamp/httpserver:${tag} .

push: release
	echo "pushing cloudnativecamp/httpserver"
	docker push cloudnativecamp/httpserver:v1.0
