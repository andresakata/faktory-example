FROM alpine:3.18
WORKDIR /usr/src/golang_app
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod vendor
