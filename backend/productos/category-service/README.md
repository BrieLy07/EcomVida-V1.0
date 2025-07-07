# 🗂️ Category Service - Ecomvida

Microservicio para la gestión de categorías y sus productos asociados en Ecomvida.

---

## 🧩 Descripción del Servicio

Permite registrar, consultar, actualizar y eliminar categorías, así como obtener los productos relacionados a una categoría. Implementa **FastAPI**, usa **PostgreSQL** y el patrón de diseño **Composite**.

---

## 🏗️ Estructura del Proyecto

category-service/
├── app/
│ ├── controllers/ # Endpoints separados
│ ├── models/ # Categoría y Producto (composición)
│ ├── repositories/ # Acceso a datos
│ ├── routes/ # Registro de rutas
│ └── db.py # Conexión a PostgreSQL
├── main.py
├── requirements.txt
├── Dockerfile
├── .env.example
└── README.md

---

## 🧪 Endpoints Disponibles

| Método | Endpoint                          | Descripción                              |
|--------|-----------------------------------|------------------------------------------|
| GET    | `/categories`                     | Listar todas las categorías              |
| GET    | `/categories/:id`                 | Obtener categoría por ID                 |
| POST   | `/categories`                     | Crear una nueva categoría                |
| PUT    | `/categories/:id`                 | Actualizar una categoría existente       |
| DELETE | `/categories/:id`                 | Eliminar una categoría                   |
| GET    | `/categories/:id/products`        | Obtener productos de una categoría       |

---

## 🧱 Patrón de Diseño Utilizado

**Composite:**  
Cada categoría puede contener múltiples productos. El patrón permite modelar esta relación jerárquica.

---

## ⚙️ Variables de Entorno

Ver `.env.example`:

PORT=3007
DB_HOST=category-db.cg9go1thnm7i.us-east-1.rds.amazonaws.com
DB_PORT=5432
DB_NAME=categories
DB_USER=postgres
DB_PASSWORD=tu_clave

---

## 🐳 Dockerización

El `Dockerfile` instala dependencias y ejecuta FastAPI en el puerto `3007`. Preparado para despliegue en EC2.

---

## ✅ Pruebas

Todos los endpoints fueron verificados usando Postman.