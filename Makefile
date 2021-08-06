.PHONY : build docker

mod-download:
	go mod download

build:
	GOOS=linux GOARCH=amd64 go build -o webhook .

docker-build:
	docker build -t webhook .