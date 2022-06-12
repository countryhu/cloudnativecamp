FROM golang:1.17 AS build
WORKDIR .
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o ./bin/amd64/homework1_httpserver main.go

FROM busybox
WORKDIR .
COPY --from=build ./bin/amd64/homework1_httpserver /homework1_httpserver
EXPOSE 8360
ENV ENV local
WORKDIR .
ENTRYPOINT ["/homework1_httpserver"]

