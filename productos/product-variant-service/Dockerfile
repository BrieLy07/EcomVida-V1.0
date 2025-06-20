FROM php:8.2-cli

WORKDIR /app

# Instalar librerías necesarias para PECL y sockets
RUN apt-get update && apt-get install -y \
    git unzip libcurl4-openssl-dev pkg-config libssl-dev \
    && docker-php-ext-install sockets

# Instalar extensión MongoDB
RUN pecl install mongodb-1.21.1 && docker-php-ext-enable mongodb

# Copiar Composer desde imagen oficial
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Copiar el proyecto
COPY . .

# Instalar dependencias PHP
RUN composer update --no-interaction


# Puerto de exposición y servidor PHP embebido
CMD ["php", "-S", "0.0.0.0:8085", "-t", "public"]
