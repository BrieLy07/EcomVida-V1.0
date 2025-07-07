FROM golang:1.21

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o warehouse ./cmd/main.go

EXPOSE 3017

CMD ["./warehouse"]
