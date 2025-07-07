# 📦 Dominio: Pedidos y Pagos - Ecomvida

Este dominio gestiona todo el ciclo de vida de un pedido: desde su creación, seguimiento, procesamiento de pagos, reembolsos y emisión de facturas.

---

## 🧱 Microservicios Incluidos

| Microservicio           | Lenguaje  | Arquitectura | Patrón Diseño     | Base de Datos       | Tipo BD    | Puerto |
|-------------------------|-----------|--------------|-------------------|----------------------|------------|--------|
| order-service           | Python    | REST         | Repository         | PostgreSQL (RDS)     | Relacional | 3010   |
| order-tracking-service  | Go        | WebSocket    | Observer           | Redis (local o nube) | Memoria    | 3011   |
| payment-service         | Python    | REST         | Template Method    | PostgreSQL (RDS)     | Relacional | 3012   |
| refund-service          | Node.js   | REST         | Command            | DynamoDB (local)     | NoSQL      | 3014   |
| invoice-service         | Node.js   | REST         | Strategy           | MongoDB Atlas        | NoSQL      | 3015   |

---

## 🧪 Pruebas

Todos los microservicios incluyen:
- Endpoints separados por archivo.
- Pruebas con Postman.
- Configuración Docker.
- Uso de `.env` y secretos protegidos.
- Listos para despliegue en EC2 o servicios compatibles del Free Tier.

---

## ☁️ Despliegue en AWS

- PostgreSQL se conecta a RDS.
- Redis y DynamoDB se simulan localmente si no se dispone de credenciales.
- MongoDB usa MongoDB Atlas Free Tier.
- Toda la arquitectura sigue el diseño optimizado para Free Tier y preparado para CI/CD.

---

## 🔁 Flujo de Trabajo

Cada microservicio fue construido y probado individualmente antes de pasar al siguiente. El dominio completo está listo para integración con el frontend y despliegue final.
