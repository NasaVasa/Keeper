FROM golang:1.20-alpine as builder

WORKDIR /keeper
COPY /bot/go.* ./
RUN go mod download

COPY /bot ./
RUN go build -o keeper.bin ./

ENTRYPOINT ["/keeper/keeper.bin"]
