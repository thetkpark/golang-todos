
FROM golang:alpine as builder
RUN apk add build-base

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o main .

FROM alpine:latest
WORKDIR /app
ENV GIN_MODE=release
COPY --from=builder /app/main ./
CMD ["/app/main"]