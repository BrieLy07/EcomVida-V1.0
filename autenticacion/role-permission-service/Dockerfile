FROM golang:1.21

WORKDIR /app

# Instalar netcat
RUN apt-get update && apt-get install -y postgresql-client


# Copiar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fuente
COPY . .

# Compilar el binario
RUN go build -o main ./cmd

# Copiar y dar permisos al script de espera
COPY wait-for-db.sh .
RUN chmod +x wait-for-db.sh

EXPOSE 8083

CMD ["./wait-for-db.sh", "./main"]
