FROM golang:alpine AS builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update && apk --no-cache add build-base
WORKDIR /go/src/github.com/linsongze/iptvproxy/
COPY . . 
RUN GOPROXY="https://goproxy.io" GO111MODULE=on go build -o iptvProxy .

FROM alpine:latest
RUN apk --no-cache add ca-certificates youtube-dl tzdata libc6-compat libgcc libstdc++
WORKDIR /root
COPY --from=builder /go/src/github.com/linsongze/iptvproxy/.env .
COPY --from=builder /go/src/github.com/linsongze/iptvproxy/iptvProxy .
EXPOSE 19000
CMD ["./iptvProxy"]
