# 🛒 Product Service - Ecomvida

Microservicio encargado de gestionar el catálogo de productos del sistema Ecomvida.

---

## 🧩 Descripción del Servicio

Permite crear, consultar, actualizar y eliminar productos. Utiliza **FastAPI** como framework web, **PostgreSQL** como base de datos y aplica el patrón de diseño **Repository** para la lógica de acceso a datos.

---

## 🏗️ Estructura del Proyecto

product-service/
├── app/
│ ├── controllers/ # Un archivo por endpoint
│ ├── models/ # Modelo Producto
│ ├── repositories/ # ProductoRepository (Patrón Repository)
│ ├── routes/ # Rutas agrupadas
│ └── db.py # Conexión a PostgreSQL
├── main.py # Inicialización FastAPI
├── requirements.txt
├── Dockerfile
├── .env.example
└── README.md

---

## 🧪 Endpoints Disponibles

| Método | Endpoint           | Descripción                    |
|--------|--------------------|--------------------------------|
| GET    | `/products`        | Obtener todos los productos    |
| GET    | `/products/:id`    | Obtener producto por ID        |
| POST   | `/products`        | Crear nuevo producto           |
| PUT    | `/products/:id`    | Actualizar producto existente  |
| DELETE | `/products/:id`    | Eliminar producto              |

---

## 🧱 Patrón de Diseño Utilizado

Se implementa el patrón **Repository** para separar la lógica de negocio del acceso a datos, encapsulando las operaciones SQL en una clase independiente (`ProductoRepository`).

---

## ⚙️ Variables de Entorno

Ver archivo `.env.example`:

PORT=3006
DB_HOST=product-db.cg9go1thnm7i.us-east-1.rds.amazonaws.com
DB_PORT=5432
DB_NAME=products
DB_USER=postgres
DB_PASSWORD=tu_clave


---

## 🐳 Dockerización

Contenedor basado en `python:3.10-slim`, instala dependencias desde `requirements.txt` y ejecuta la app con `uvicorn`.

---

## ✅ Pruebas

Todos los endpoints fueron validados usando Postman.