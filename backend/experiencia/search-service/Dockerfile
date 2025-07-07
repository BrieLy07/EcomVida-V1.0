# Etapa 1: Build
FROM rust:1.77 as builder

WORKDIR /usr/src/app

# Copiar solo los archivos necesarios para cache eficiente
COPY Cargo.toml Cargo.lock ./
RUN mkdir src && echo "fn main() {}" > src/main.rs
RUN cargo build --release
RUN rm -rf src

# Copiar el código fuente
COPY . .
RUN cargo build --release

# Etapa 2: Imagen final
FROM debian:bullseye-slim

# Instalar dependencias necesarias
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Crear usuario no root
RUN useradd -m appuser
USER appuser

# Copiar binario desde la etapa anterior
COPY --from=builder /usr/src/app/target/release/search-service /usr/local/bin/search-service

EXPOSE 3030

CMD ["search-service"]
