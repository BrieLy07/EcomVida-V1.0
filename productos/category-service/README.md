# Category Service - ECOMVIDA

Microservicio REST en Node.js para gestión de categorías de productos.

---

## 📦 Endpoints

| Método | Ruta                  | Descripción                 |
|--------|-----------------------|-----------------------------|
| POST   | `/api/categorias`     | Crear nueva categoría       |
| GET    | `/api/categorias`     | Listar todas las categorías |
| GET    | `/api/categorias/:id` | Buscar categoría por ID     |

---

## ⚙️ Variables `.env`

```env
DB_HOST=category-db
DB_PORT=3306
DB_USER=root
DB_PASSWORD=admin123
DB_NAME=categorydb
PORT=3002
