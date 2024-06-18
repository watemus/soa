FROM golang:1.22.4-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o auth-service ./cmd/main.go

CMD ["./auth-service"]
