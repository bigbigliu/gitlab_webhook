FROM golang:latest AS builder

WORKDIR /go/cache

# 利用docker镜像文件分层做缓存
ADD go.mod .
ADD go.sum .
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download

# 设置工作目录
WORKDIR /src

# 将当前目录的go工程代码复制到docker容器工作目录下
COPY . /src

# GOOS:目标系统为linux 
# CGO_ENABLED:默认为1，启用C语言版本的GO编译器，通过设置成0禁用它
# GOARCH:32位系统为386，64位系统为amd64
# -o:指定编译后的可执行文件名称
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o app

# 运行：使用scratch作为基础镜像
#FROM scratch as prod
FROM alpine as prod

RUN  echo "Asia/Shanghai" > /etc/timezone \
    && rm -f /etc/localtime \
    && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 复制 conf 文件到容器 conf
COPY --from=builder /src/conf/ . /svr/

# Expose port 8080 to the outside world
EXPOSE 8085

# Command to run the executable
CMD ["./app"]



