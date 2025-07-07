# 📦 Order Service - Ecomvida

Microservicio REST encargado de gestionar los pedidos realizados por los usuarios. Implementado en Python con FastAPI y PostgreSQL. Aplica el patrón de diseño **Repository**.

---

## 🧩 Descripción del Servicio

Permite crear, consultar, actualizar y eliminar pedidos del sistema. Toda la lógica de acceso a datos está desacoplada mediante un repositorio.

---

## 🏗️ Estructura del Proyecto

order-service/
├── app/
│ ├── controllers/ # Un archivo por endpoint
│ ├── database/ # Conexión a PostgreSQL
│ ├── models/ # Modelo de datos Pedido
│ └── repositories/ # Lógica de acceso (Repository)
├── main.py # Inicializa FastAPI
├── requirements.txt # Librerías necesarias
├── Dockerfile # Imagen optimizada
├── .env.example # Variables de entorno
└── README.md # Este archivo

---

## 🔁 Endpoints REST

| Método | Endpoint         | Descripción                  |
|--------|------------------|------------------------------|
| GET    | `/orders`        | Listar todos los pedidos     |
| GET    | `/orders/:id`    | Obtener pedido por ID        |
| POST   | `/orders`        | Crear nuevo pedido           |
| PUT    | `/orders/:id`    | Actualizar pedido existente  |
| DELETE | `/orders/:id`    | Eliminar o cancelar pedido   |

---

## 🧠 Patrón de Diseño

**Repository**  
Toda la lógica de base de datos está aislada en `repositories/order_repository.py`, facilitando pruebas y mantenimiento.

---

## ⚙️ Variables de Entorno

Ver `.env.example`:

PORT=3011
DB_HOST=order-db.cg9go1thnm7i.us-east-1.rds.amazonaws.com
DB_PORT=5432
DB_NAME=orders
DB_USER=admin
DB_PASSWORD=admin123

## ⚙️ Dockerización

El contenedor expone el puerto 3011 y puede integrarse a un docker-compose general.



