FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download 
RUN go build -o sales-app ./cmd/main.go

CMD ["./sales-app"]