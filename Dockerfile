FROM golang:latest

ENV GO111MODULE=on
ENV ServerAddress=8080
ENV ContextTimeout=15
ENV DB_HOST=mongodb://user:pass@mongo:27018/mezink
ENV DB_PORT=3306
ENV DB_USER=user
ENV DB_PASS=pass
ENV DB_NAME=mezink

WORKDIR /app
COPY go.mod /app
COPY go.sum /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
COPY . /app
ENTRYPOINT CompileDaemon --build="go build -o main" --command=./cmd/main