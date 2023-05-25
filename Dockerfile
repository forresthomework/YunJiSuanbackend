# 使用 Golang 官方提供的 Golang 镜像作为基础镜像
FROM golang:1.19

# 设置工作目录
WORKDIR /work

# 复制整个项目到工作目录
COPY . .

# 复制 go.mod 和 go.sum 文件到工作目录
RUN go mod download

# 执行 test.go 中的 test_1 函数
RUN go test -run Test_Convet_TXT_2_Redis
# 执行 main.go
expose 9999
CMD ["go", "run", "main/main.go"]

