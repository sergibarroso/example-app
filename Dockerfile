# syntax=docker/dockerfile:1

### Build app
FROM golang:1.20.5-alpine as build

WORKDIR /build
COPY . /build

RUN go mod download && go mod verify
RUN go build -v -o example-app ./...

### Generate user
FROM alpine:3 as user

WORKDIR /
RUN echo "celonis:x:1000:1000:,,,:/app:/bin/nologin" >passwd
RUN echo "celonis:x:1000:celonis" >group

### Live image
FROM scratch as live

LABEL group="cloud.celonis"
LABEL version="1.1.1"

COPY --from=user /passwd /etc/passwd
COPY --from=user /group /etc/group

USER celonis

WORKDIR /app
COPY assets /app/assets
COPY --from=build /build/example-app .

EXPOSE 8080/tcp
ENTRYPOINT ["./example-app"]
