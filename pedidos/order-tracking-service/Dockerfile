FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o tracking-app main.go

EXPOSE 8088

CMD ["./tracking-app"]
