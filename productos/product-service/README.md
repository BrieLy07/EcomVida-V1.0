# Product Service - ECOMVIDA

Microservicio REST en Python (FastAPI) que gestiona productos. Utiliza PostgreSQL como base de datos y SQLAlchemy como ORM.

---

## 📦 Endpoints

| Método | Ruta                    | Descripción                  |
|--------|-------------------------|------------------------------|
| POST   | `/productos/`           | Crear un nuevo producto      |
| GET    | `/productos/`           | Listar todos los productos   |
| GET    | `/productos/{id}`       | Obtener producto por ID      |

---

## 🧱 Tecnologías

- Python 3.11
- FastAPI
- PostgreSQL 14
- SQLAlchemy
- Docker & Docker Compose

---

## ⚙️ Variables de entorno (`.env`)

```env
DB_HOST=product-db
DB_PORT=5432
DB_NAME=productsdb
DB_USER=postgres
DB_PASSWORD=admin123
