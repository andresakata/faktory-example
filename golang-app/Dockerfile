FROM golang:1.20

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app main/main.go
RUN go build -v -o /usr/local/bin/enqueue enqueue/enqueue.go

CMD ["app"]
