
 FROM golang:1.17.1-alpine
 
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn \
    && echo export PATH="$PATH:$(go env GOPATH)/bin" >> ~/.bashrc


# 创建文件夹
RUN mkdir /app
 
# 设置工作目录

#将Dockerfile 中的文件存储到/app下
ADD . /app

# ENV GO111MODULE=on
# ENV GOPROXY="https://goproxy.io"

WORKDIR /app

RUN go mod download && go mod verify
# 因为已经是在 /app下了，所以使用  ./
RUN go build -o main ./JoinV2/main/mainRecv/mainRecv.go
 
# 暴露的端口
EXPOSE 50051
 
#设置容器的启动命令，CMD是设置容器的启动指令
CMD /app/main
 