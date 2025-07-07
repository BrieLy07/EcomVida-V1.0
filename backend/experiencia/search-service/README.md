# Search-Service - Ecomvida

Microservicio de búsqueda para productos en Ecomvida.

## 📦 Descripción

Este microservicio permite buscar productos por texto, categoría o marca, y obtener sugerencias en tiempo real. Usa Elasticsearch como motor de búsqueda y Redis como sistema de cache (patrón Cache Proxy).

## 🔧 Tecnología

- Lenguaje: Rust
- Arquitectura: REST
- Patrón de diseño: Cache Proxy
- Base de datos: Elasticsearch (otros)
- Almacenamiento en memoria: Redis

## 📁 Estructura del Proyecto

search-service/
├── src/
│ ├── handlers/ # Lógica de endpoints
│ ├── elastic/ # Conexión y funciones con Elasticsearch
│ ├── cache/ # Cache Redis
│ ├── routes.rs # Rutas del microservicio
│ └── main.rs # Punto de entrada
├── .env
├── .gitignore
├── Dockerfile
└── docker-compose.yml

## 🧪 Endpoints Disponibles

| Método | Endpoint                  | Descripción                                      |
|--------|---------------------------|--------------------------------------------------|
| GET    | `/search`                 | Buscar productos                                 |
| GET    | `/search/suggestions`     | Obtener sugerencias de búsqueda en tiempo real   |

## 🚀 Ejecución

```bash
docker-compose up --build
```

## 🧠 Notas
Redis cachea cada búsqueda por 60 segundos.

Elasticsearch consulta el índice products.


