FROM golang:1.17.8 AS build

WORKDIR /app

ADD ./*  /app/
RUN GOPROXY=https://proxy.golang.com.cn,direct  go mod download
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o /app/hello-dapr

FROM alpine:3.15.3
WORKDIR /app
COPY --from=build /app/hello-dapr /app/hello-dapr
ENTRYPOINT ["/app/hello-dapr"]
