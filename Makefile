.PHONY : build docker

mod-download: # 下载依赖
	go mod download

build:
	GOOS=linux GOARCH=amd64 go build -o webhook .

docker-build:
	docker build -t webhook .