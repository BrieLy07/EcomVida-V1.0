# User Profile Service - ECOMVIDA

Microservicio para gestionar el perfil de usuario autenticado dentro de la plataforma ECOMVIDA. Funciona bajo arquitectura GraphQL, conectado a MongoDB y protegido por token JWT emitido desde el `auth-service`.

---

## 🧱 Tecnologías
- Python 3.11
- Flask + GraphQL
- MongoDB
- JWT
- Docker & Docker Compose

---

## 📦 Operaciones disponibles

### Consulta
```graphql
query {
  getPerfil {
    nombre
    apellido
    correo
    numero
    foto
  }
}

---

### Mutación

mutation {
  actualizarPerfil(
    nombre: "Juan",
    apellido: "Pérez",
    correo: "juan@correo.com",
    numero: "0999999999",
    foto: "url.jpg"
  ) {
    ok
  }
}

### Autorización

Todas las operaciones requieren un Header
Authorization: Bearer <TOKEN_JWT>

### ⚙️ Variables de entorno

Crea un archivo .env basado en:

env
Copiar
Editar
PORT=5000
MONGO_URI=mongodb://mongo:27017
DB_NAME=perfilusuario
JWT_SECRET=miclavesecreta

### Levantar con Docker

docker compose up --build
http://localhost:5000/graphql
