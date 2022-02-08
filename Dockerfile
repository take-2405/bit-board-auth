
FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/take-2405/bit-board-auth
COPY . .
RUN go build -o server cmd/main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/miraikeitai2020/backend-record/server /app

CMD /app $PORT