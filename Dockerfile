FROM golang:latest AS build
RUN go get -v -u 'github.com/prometheus/client_golang/prometheus'

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -o main -i main.go

FROM ubuntu:18.04 AS release
MAINTAINER Alex Sirmais
ENV TZ=Europe/Moscow

USER root
WORKDIR /app
COPY --from=build /app/main .
RUN chmod +x ./main


EXPOSE 8080/tcp


CMD ./main