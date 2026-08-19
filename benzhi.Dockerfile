# 评测专用：保留完整 Go 工具链（勿改成多阶段/只留二进制）
FROM golang:1.22

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build ./...

CMD ["bash"]
