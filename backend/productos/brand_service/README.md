# 🏷️ Brand Service - Ecomvida

Microservicio encargado de gestionar las marcas de productos y sus productos asociados en Ecomvida.

---

## 🧩 Descripción del Servicio

Permite crear, consultar, actualizar y eliminar marcas, así como obtener los productos de una marca específica. Utiliza **FastAPI**, **PostgreSQL**, y aplica el patrón de diseño **Adapter** para transformar las filas de productos.

---

## 🏗️ Estructura del Proyecto

brand-service/
├── app/
│ ├── controllers/ # Endpoints separados
│ ├── models/ # Marca y Producto
│ ├── adapters/ # Adaptador para productos
│ ├── repositories/ # MarcaRepository
│ ├── routes/ # Agrupación de rutas
│ └── db.py # Conexión PostgreSQL
├── main.py
├── requirements.txt
├── Dockerfile
├── .env.example
└── README.md

---

## 🧪 Endpoints Disponibles

| Método | Endpoint                      | Descripción                             |
|--------|-------------------------------|-----------------------------------------|
| GET    | `/brands`                     | Listar todas las marcas                 |
| GET    | `/brands/:id`                 | Obtener marca por ID                    |
| POST   | `/brands`                     | Crear una nueva marca                   |
| PUT    | `/brands/:id`                 | Actualizar una marca                    |
| DELETE | `/brands/:id`                 | Eliminar una marca                      |
| GET    | `/brands/:id/products`        | Obtener productos de una marca          |

---

## 🧱 Patrón de Diseño Utilizado

**Adapter:**  
Se utiliza para transformar los datos de productos antes de devolverlos en la respuesta, desacoplando la estructura interna de la tabla de productos.

---

## ⚙️ Variables de Entorno

Ver `.env.example`:

PORT=3008
DB_HOST=brand-db.cg9go1thnm7i.us-east-1.rds.amazonaws.com
DB_PORT=5432
DB_NAME=brands
DB_USER=postgres
DB_PASSWORD=tu_clave

---

## 🐳 Dockerización

El servicio se ejecuta en el puerto `3008` utilizando Uvicorn. Preparado para despliegue en EC2 o entorno local.

---

## ✅ Pruebas

Todos los endpoints fueron validados utilizando Postman.