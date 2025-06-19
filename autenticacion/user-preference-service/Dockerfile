FROM php:8.2-cli

WORKDIR /app

# Instalar dependencias necesarias para AWS SDK
RUN apt-get update && apt-get install -y unzip git libzip-dev \
    && docker-php-ext-install zip

# Instalar Composer
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

COPY . .

RUN composer install

EXPOSE 80
CMD ["php", "-S", "0.0.0.0:80", "-t", "src"]
