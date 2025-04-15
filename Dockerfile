FROM golang:1.23-alpine AS builder

WORKDIR /app

# 安装时区数据（构建阶段）
RUN apk update && apk add --no-cache tzdata

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    TZ=Asia/Shanghai

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN go build -o al_learn_ass .

# 使用轻量级的 alpine 作为最终镜像
FROM alpine:latest

# 在最终镜像中也安装 tzdata！
RUN apk update && apk add --no-cache tzdata

# 设置时区环境变量（覆盖 builder 阶段的设置）
ENV TZ=Asia/Shanghai

# 复制时区文件到系统目录（可选但更可靠）
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app
COPY --from=builder /app/al_learn_ass .
COPY --from=builder /app/view/pages /app/view/pages
COPY --from=builder /app/config.yaml .

ENTRYPOINT ["/app/al_learn_ass"]
