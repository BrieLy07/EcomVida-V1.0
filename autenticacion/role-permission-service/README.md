# Role Permission Service - ECOMVIDA

Microservicio REST en Go que gestiona roles y permisos de usuarios. Utiliza JWT para autenticación y PostgreSQL para almacenamiento.

---

## 🚀 Funcionalidades
- Crear roles
- Asignar roles a usuarios
- Consultar roles del usuario autenticado

---

## 🛠 Tecnologías
- Go 1.21
- PostgreSQL 14
- Gorilla Mux
- JWT
- Docker & Docker Compose

---

## 🔧 Variables de entorno (.env)

```env
PORT=8083
DB_HOST=role-db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=admin123
DB_NAME=rolesdb
JWT_SECRET=miclavesecreta

```

### 📦 Endpoints

| Método | Ruta                 | Autenticado | Descripción              |
| ------ | -------------------- | ----------- | ------------------------ |
| POST   | `/api/roles/`        | ✅           | Crear un nuevo rol       |
| POST   | `/api/roles/asignar` | ✅           | Asignar rol a un usuario |
| GET    | `/api/roles/`        | ✅           | Ver roles del usuario    |


### 🧪 Pruebas en Postman

Todas las rutas requieren:

Copiar
Editar
Authorization: Bearer <TOKEN_JWT_DEL_AUTH_SERVICE>

### 🐳 Docker

docker compose up --build

Servidor: http://localhost:8083