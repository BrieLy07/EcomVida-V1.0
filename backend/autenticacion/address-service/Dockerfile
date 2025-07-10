# Etapa base
FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o address-service .

EXPOSE 3003

CMD ["./address-service"]
