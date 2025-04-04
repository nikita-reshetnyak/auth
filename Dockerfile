FROM golang:1.24.1-alpine AS builder

COPY . /github.com/nikita-reshetnyak/auth/source
WORKDIR /github.com/nikita-reshetnyak/auth/source

RUN go mod download
RUN go build -o ./bin/crud_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/nikita-reshetnyak/auth/source/bin/crud_server .

CMD ["./crud_server"]