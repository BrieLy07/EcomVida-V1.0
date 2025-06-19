#!/bin/sh

echo "Esperando a que PostgreSQL esté completamente disponible..."

until pg_isready -h $DB_HOST -p $DB_PORT -U $DB_USER > /dev/null 2>&1; do
  sleep 1
done

echo "PostgreSQL listo. Iniciando servicio..."
exec "$@"
