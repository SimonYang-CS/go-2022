FROM golang:1.17.8 AS build

WORKDIR /app

ADD ./*  /app/
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o /app/simple-error


FROM alpine:3.15.3
WORKDIR /app
COPY --from=build /app/simple-error /app/simple-error
ENTRYPOINT ["/app/simple-error"]
