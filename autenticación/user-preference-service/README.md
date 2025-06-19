# User Preferences Service - ECOMVIDA

Microservicio en PHP que permite guardar y consultar las preferencias personales de un usuario autenticado. Funciona mediante WebHook y DynamoDB local, con autenticación vía JWT.

---

## 🧱 Tecnologías
- PHP 8.2
- AWS DynamoDB Local
- JWT (validación)
- Docker & Docker Compose
- AWS SDK PHP

---

## 📦 Endpoints disponibles

| Método | Ruta              | Descripción                           |
|--------|-------------------|---------------------------------------|
| POST   | `/webhook`        | Guarda o actualiza preferencias       |
| GET    | `/preferencias`   | Retorna las preferencias del usuario  |

---

## ⚙️ Variables de entorno

```env
AWS_REGION=us-west-2
AWS_ACCESS_KEY_ID=test
AWS_SECRET_ACCESS_KEY=test
DYNAMODB_ENDPOINT=http://dynamodb-local:8000
DYNAMODB_TABLE=preferencias_usuario
JWT_SECRET=miclavesecreta

```
---

## 🧪 Pruebas con Postman

POST /webhook
json
Copiar
Editar
{
  "idioma": "es",
  "tema": "oscuro",
  "favoritos": ["tecnología", "hogar"]
}

GET /preferencias
Ambos endpoints requieren:

makefile
Copiar
Editar
Authorization: Bearer <TOKEN_DEL_AUTH_SERVICE>

---


