FROM golang:1.20-bullseye

ENV GO111MODULE=on
ENV ServerAddress=8080
ENV ContextTimeout=15
ENV DB_HOST=mongodb://mongo:27017/mezink
ENV DB_PORT=27017
ENV DB_USER=user
ENV DB_PASS=pass
ENV DB_NAME=mezink

WORKDIR /app
COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app
COPY ./cmd/main.go /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o /test-go-mezink
EXPOSE 8080
CMD ["/test-go-mezink"]