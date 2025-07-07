# 📦 Order Tracking Service - Ecomvida

Microservicio de seguimiento de pedidos en tiempo real. Implementado en Go usando WebSocket, REST y Redis. Aplica el patrón de diseño **Observer**.

---

## 🧩 Descripción del Servicio

Permite consultar y actualizar el estado de un pedido, así como suscribirse a eventos en tiempo real mediante WebSocket para recibir notificaciones de cambios de estado.

---

## 🏗️ Estructura del Proyecto

order-tracking-service/
├── app/
│ ├── redis/ # Cliente Redis + lógica Observer
│ ├── rest/ # Endpoints REST
│ └── ws/ # WebSocket tracking en tiempo real
├── main.go # Entrada del programa
├── Dockerfile # Imagen optimizada
├── .env.example # Variables de entorno
├── go.mod # Dependencias Go
└── README.md # Este archivo

---

## 🔁 Endpoints

| Método      | Endpoint                    | Descripción                            |
|-------------|-----------------------------|----------------------------------------|
| GET         | `/orders/:id/tracking`      | Obtener el estado actual de un pedido |
| POST        | `/orders/:id/tracking`      | Crear o actualizar estado de pedido   |
| WebSocket   | `/ws/orders/:id/track`      | Canal en tiempo real por pedido       |

---

## 🧠 Patrón de Diseño

**Observer**  
Redis actúa como intermediario publicador/subscriptor entre los productores de cambios y los clientes WebSocket.

---

## ⚙️ Variables de Entorno

PORT=3012
REDIS_HOST=redis
REDIS_PORT=6379

## 🐳 Dockerización

El contenedor expone el puerto 3012. Redis debe estar corriendo en el mismo entorno o conectado externamente.
