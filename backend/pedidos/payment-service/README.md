# 💳 Payment Service - Ecomvida

Microservicio para procesar pagos asociados a pedidos. Implementado en Python con FastAPI y PostgreSQL. Aplica el patrón de diseño **Template Method**.

---

## 🧩 Descripción del Servicio

Permite registrar pagos, consultar un pago por ID y obtener todos los pagos asociados a un pedido.

---

## 🏗️ Estructura del Proyecto

payment-service/
├── app/
│ ├── controllers/ # Endpoints (uno por archivo)
│ ├── database/ # Conexión a PostgreSQL
│ ├── models/ # Modelo Pago
│ └── templates/
│ └── procesadores/ # Implementación Template Method
├── main.py # Configuración FastAPI
├── Dockerfile # Imagen optimizada
├── .env.example # Variables de entorno
├── requirements.txt # Dependencias
└── README.md

---

## 🔁 Endpoints

| Método | Endpoint                          | Descripción                         |
|--------|-----------------------------------|-------------------------------------|
| POST   | `/payments`                       | Procesar nuevo pago                 |
| GET    | `/payments/{id}`                  | Obtener pago por ID                 |
| GET    | `/payments/order/{order_id}`      | Obtener pagos por ID de pedido      |

---

## 🧠 Patrón de Diseño

**Template Method**  
Define un flujo de pasos para procesar un pago. El método base orquesta validación, ejecución y persistencia. Las subclases concretan los detalles por método de pago.

---

## ⚙️ Variables de Entorno

```env
PORT=3013
DB_HOST=payment-db.cg9go1thnm7i.us-east-1.rds.amazonaws.com
DB_PORT=5432
DB_NAME=ecomvida
DB_USER=admin
DB_PASSWORD=admin123