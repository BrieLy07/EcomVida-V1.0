# Auth Service - ECOMVIDA

Microservicio de autenticación para la plataforma eCommerce ECOMVIDA. Permite el registro y login de usuarios mediante API REST y manejo de JWT.
..

---

## 🧱 Tecnologías
- Node.js
- Express.js
- PostgreSQL
- Sequelize ORM
- JWT (autenticación)
- Docker & Docker Compose

---

## 📦 Endpoints disponibles

| Método | Ruta                   | Descripción                |
|--------|------------------------|----------------------------|
| POST   | `/api/auth/registro`   | Registro de nuevo usuario |
| POST   | `/api/auth/login`      | Inicio de sesión           |

---

## 📁 Estructura del proyecto

auth-service/
├── frontend/ # HTML para pruebas visuales
├── src/ # Código principal
│ ├── config/ # Conexión base de datos
│ ├── controllers/ # Lógica de negocio
│ ├── models/ # Modelo Usuario
│ ├── routes/ # Rutas de la API
│ └── index.js # Servidor principal
├── Dockerfile
├── docker-compose.yml
├── .env
├── .env.example
└── README.md

---

## ⚙️ Variables de entorno

Crea un archivo `.env` basado en el siguiente ejemplo:

```env
PORT=3000
DB_HOST=auth-db
DB_PORT=5432
DB_NAME=authdb
DB_USER=postgres
DB_PASSWORD=admin123
JWT_SECRET=miclavesecreta

---

## ⚙️ Levantar con Docker

- docker compose up --build

- http://localhost:3000/api/auth

## 🧪 Pruebas

Puedes utilizar Postman o el frontend HTML disponible en frontend/ para probar el registro y login de usuarios.

## 🛠 CI con GitHub Actions

Este microservicio cuenta con un flujo CI que valida dependencias, build y levantamiento básico en cada push a la rama.

