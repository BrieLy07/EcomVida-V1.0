# recommendation-service - Ecomvida

Microservicio gRPC para generar recomendaciones personalizadas y relacionadas de productos, así como recibir retroalimentación del usuario.

## 📦 Estructura

- `app/handlers`: Servidores gRPC que responden a las llamadas.
- `app/chain`: Implementación del patrón Chain of Responsibility.
- `app/database`: Conexión a PostgreSQL y funciones CRUD.
- `app/server.py`: Servidor gRPC principal.
- `proto/recommendation.proto`: Especificación de servicios y mensajes gRPC.

## 🧩 Endpoints (vía gRPC)

| RPC                      | Descripción                                         |
|--------------------------|-----------------------------------------------------|
| `GetRecommendations`     | Obtener recomendaciones personalizadas por usuario |
| `GetRelatedProducts`     | Obtener productos relacionados                      |
| `SendFeedback`           | Registrar feedback del usuario                     |

## 🧠 Patrón de diseño

- **Chain of Responsibility**: Para encadenar el flujo entre recomendaciones personalizadas y relacionadas.
- **Separación de responsabilidades**: cada handler y función encapsula lógica concreta.

## 🧪 Base de Datos

- **PostgreSQL**
- Tablas: `recommendations`, `related_products`, `feedback`

## ⚙️ Variables de Entorno (.env)

```
APP_PORT=3005
DB_HOST=...
DB_PORT=5432
DB_USER=...
DB_PASSWORD=...
DB_NAME=...
```

## 🚀 Despliegue

Este servicio debe desplegarse en una instancia EC2 usando PostgreSQL en RDS. Las credenciales estarán protegidas usando `.env` y secretos en GitHub Actions.

