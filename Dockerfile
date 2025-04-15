FROM golang:1.23-alpine AS builder

WORKDIR /app

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

WORKDIR /app
COPY --from=builder --chown=runner /app/al_learn_ass .
COPY --chown=runner config.yaml .

ENTRYPOINT ["/app/al_learn_ass"]