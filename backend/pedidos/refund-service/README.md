# 💸 Refund Service - Ecomvida

Microservicio para gestionar solicitudes de reembolsos asociados a pedidos. Implementado en Node.js con arquitectura REST, usa **DynamoDB local** y el patrón de diseño **Command**.

---

## 🧩 Descripción del Servicio

Permite crear reembolsos, obtener uno por ID y listar todos los asociados a un pedido.

---

## 🏗️ Estructura del Proyecto

refund-service/
├── app/
│ ├── commands/ # Comandos (crear, obtener, listar)
│ ├── controllers/ # Endpoints (uno por archivo)
│ └── database/ # Cliente DynamoDB
├── main.js # Inicializa Express
├── Dockerfile # Imagen optimizada
├── .env.example # Variables de entorno
├── package.json # Dependencias
└── README.md

---

## 🔁 Endpoints

| Método | Endpoint                          | Descripción                         |
|--------|-----------------------------------|-------------------------------------|
| POST   | `/refunds`                        | Solicitar reembolso                 |
| GET    | `/refunds/:id`                    | Obtener detalles de un reembolso    |
| GET    | `/refunds/order/:orderId`         | Reembolsos asociados a un pedido    |

---

## 🧠 Patrón de Diseño

**Command**  
Cada acción (crear, obtener, listar) está encapsulada en su propia clase con un método `ejecutar()`.

---

## ⚙️ Variables de Entorno

```env
PORT=3014
DYNAMODB_ENDPOINT=http://dynamodb-local:8000
DYNAMO_TABLE=reembolsos
```

## 📦 Dependencias

```
express
dotenv
uuid
@aws-sdk/client-dynamodb
@aws-sdk/lib-dynamodb

```

## 🐳 Docker Compose

