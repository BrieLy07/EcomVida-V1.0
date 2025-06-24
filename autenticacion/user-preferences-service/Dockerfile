FROM php:8.1-cli

RUN apt-get update && apt-get install -y unzip git \
    && docker-php-ext-install sockets

COPY . /var/www/html
WORKDIR /var/www/html

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer \
    && composer install

EXPOSE 3003

CMD ["php", "-S", "0.0.0.0:3003", "index.php"]
