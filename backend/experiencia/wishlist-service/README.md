# wishlist-service

Microservicio encargado de gestionar la lista de deseos de los usuarios en Ecomvida.

---

## 📦 Estructura del Proyecto

wishlist-service/
├── controllers/
├── models/
├── observers/
├── routes/
├── utils/
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── .gitignore
├── app.js
└── package.json


---

## 🚀 Endpoints REST

| Método | Endpoint               | Descripción                                      |
|--------|------------------------|--------------------------------------------------|
| GET    | `/wishlist`            | Obtener la lista de deseos del usuario actual   |
| POST   | `/wishlist`            | Agregar un producto a la lista de deseos        |
| DELETE | `/wishlist/:productId` | Eliminar un producto de la lista de deseos      |

> ⚠️ Requiere encabezado `x-user-id` para identificar al usuario.

---

## 🧩 Patrón de Diseño

Este microservicio aplica el patrón **Observer** para notificar cada acción realizada sobre la lista de deseos.

---

### Variables de entorno (`.env.example`)

```env
AWS_REGION=us-east-1
DYNAMO_ENDPOINT=http://dynamodb-local:8000
DYNAMO_TABLE=WISHLIST_TABLE
PORT=3026
```




