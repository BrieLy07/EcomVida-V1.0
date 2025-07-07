# Delivery Service - Ecomvida

Microservicio responsable de gestionar las reseñas de productos dentro de la plataforma Ecomvida. Utiliza GraphQL para las operaciones y se conecta a MongoDB como base de datos NoSQL.

---

## 📦 Estructura del Proyecto

review-service/
├── app/
│ ├── database/
│ │ └── connection.py
│ ├── decorators/
│ │ └── log_decorator.py
│ ├── graphql/
│ │ ├── queries/
│ │ │ └── review_queries.py
│ │ ├── mutations/
│ │ │ └── review_mutations.py
│ │ └── schema.py
│ ├── models/
│ │ └── review_model.py
├── main.py
├── requirements.txt
├── Dockerfile
├── .env.example
└── .gitignore


---

## 🚀 Endpoints GraphQL

**URL base:** `/graphql`

### 📥 Queries

| Nombre                 | Descripción                             |
|------------------------|-----------------------------------------|
| `obtenerResenas`       | Listar todas las reseñas                |
| `resenaPorId(id)`      | Obtener una reseña por ID               |
| `resenasPorProducto(producto_id)` | Obtener reseñas de un producto     |

### ✏️ Mutations

| Nombre                     | Descripción                             |
|----------------------------|-----------------------------------------|
| `crearResena(...)`         | Crear una nueva reseña                  |
| `editarResena(id, ...)`    | Editar una reseña existente             |
| `eliminarResena(id)`       | Eliminar una reseña por ID              |

---

## 🧩 Patrón de Diseño Utilizado

Este microservicio aplica el patrón **Decorator** para registrar operaciones y permitir validaciones reutilizables.

---

## 🐳 Docker

### Dockerfile

Empaqueta el microservicio para producción con una imagen liviana basada en `python:3.11-slim`.

### Variables de Entorno (`.env`)

```env
PORT=3025
MONGO_URI=mongodb+srv://usuario:contraseña@cluster.mongodb.net/ecomvida
DB_NAME=ecomvida
```