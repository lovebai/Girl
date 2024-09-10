# 第一阶段：构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制源代码
COPY . .

RUN apk add --no-cache gcc musl-dev sqlite-dev

ENV CGO_ENABLED=1

# 编译应用
RUN go build -o girl .

# 第二阶段：运行阶段
FROM alpine:latest

# 设置工作目录
WORKDIR /Girl

# 复制编译好的二进制文件
COPY --from=builder /app/girl ./girl
COPY --from=builder /app/data/ ./data

# 安装时区数据
RUN apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && chmod +x ./girl 

# 暴露端口
EXPOSE 5200

# 创建数据卷
VOLUME [ "/Girl/data"]

# 设置入口点
ENTRYPOINT ["./girl"]
