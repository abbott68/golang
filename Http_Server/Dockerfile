# 使用官方的 Golang 镜像作为基础映像
FROM golang:1.17

# 设置工作目录
WORKDIR /app

# 将应用程序的源代码复制到容器中的工作目录
COPY . .

# 构建应用程序
RUN go build -o myapp

# 设置容器启动时要运行的命令
CMD ["./myapp"]
