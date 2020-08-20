# Builder
FROM golang:1.14-alpine as builder

RUN apk update && apk upgrade && \
    apk --no-cache --update add git make

WORKDIR /go/src/github.com/dwadp/auth-example

COPY . .

RUN go get -u && \
    go build main.go

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --no-cache --update add ca-certificates tzdata && \
    mkdir auth-example && mkdir ./auth-example/config

WORKDIR /auth-example

EXPOSE 5000

ENV APP_ENV=production

# Set GIN_MODE to "release" so that the application do not print a debuggable message
ENV GIN_MODE=release

COPY --from=builder /go/src/github.com/dwadp/auth-example/main .
COPY --from=builder /go/src/github.com/dwadp/auth-example/config/production.yaml ./config/

CMD ["./main"]