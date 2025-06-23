# order-service

Microservicio de pedidos para EcomVida. Permite registrar y consultar órdenes de compra de los usuarios.

## 🔧 Tecnologías

- Java 17 + Spring Boot
- PostgreSQL
- Docker & Docker Compose
- GitHub Actions CI

## 🚀 Endpoints

### Crear orden

POST /ordenes

Body:
```json
{
  "usuario": "cliente01",
  "total": 499.99,
  "estado": "pendiente"
}

```

## Listar ordenes

GET /ordenes
