FROM golang:1.11-alpine3.8 AS build
RUN echo http://mirror.yandex.ru/mirrors/alpine/v3.5/main > /etc/apk/repositories; \
  echo http://mirror.yandex.ru/mirrors/alpine/v3.5/community >> /etc/apk/repositories
# เอาไว้ตอน link ของ golang image มันเสีย
# https://github.com/gliderlabs/docker-alpine/issues/279
RUN apk update \
  RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go
COPY ./main.go ./main.go
RUN go get -u github.com/lib/pq github.com/labstack/echo github.com/manyminds/api2go/jsonapi github.com/tidwall/gjson github.com/dgrijalva/jwt-go github.com/joho/godotenv cloud.google.com/go/profiler
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/lvgo ./main.go

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/bin /go/bin
COPY ./build ./build
#EXPOSE 8080
#ENTRYPOINT /go/bin/lvgo --port 9000
EXPOSE 8080
ENTRYPOINT /go/bin/lvgo --port 1234