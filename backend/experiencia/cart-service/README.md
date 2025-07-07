# 🛒 Cart-Service - Ecomvida

Microservicio del carrito de compras del usuario para Ecomvida.

---

## 🧱 Estructura del Proyecto

cart-service/
├── controllers/
│ ├── add_to_cart.go
│ ├── clear_cart.go
│ ├── get_cart.go
│ ├── remove_cart_item.go
│ └── update_cart_item.go
├── memento/
│ └── cart_memento.go
├── redis/
│ └── redis_client.go
├── routes/
│ └── cart_routes.go
├── .gitignore
├── go.mod
├── go.sum
└── main.go



---

## 🌐 Endpoints Disponibles

| Método | Endpoint               | Descripción                              |
|--------|------------------------|------------------------------------------|
| GET    | `/cart`                | Obtener el carrito actual del usuario    |
| POST   | `/cart`                | Agregar un producto al carrito           |
| PUT    | `/cart/:productId`     | Actualizar cantidad de un producto       |
| DELETE | `/cart/:productId`     | Eliminar un producto del carrito         |
| DELETE | `/cart/clear`          | Vaciar todo el carrito                   |

> ⚠️ Todos los endpoints requieren el header `X-User-Id` con el ID del usuario.

---

## 🧩 Patrón de Diseño Utilizado

**Memento**  
Permite capturar y restaurar el estado del carrito en Redis.

---

## 🐘 Base de Datos

- **Redis** (almacenamiento en memoria)
- Uso de claves por `usuario_id` para persistencia por sesión

---

## 🐳 Docker

Este microservicio está listo para ser dockerizado y desplegado. Redis puede ejecutarse local o remotamente según tu configuración de infraestructura.

---


