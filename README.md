# 🛒 EcomVida – Plataforma eCommerce por Microservicios

Proyecto modular de eCommerce basado en arquitectura de microservicios. Cada dominio está desacoplado tecnológicamente y se despliega de forma independiente sobre infraestructura EC2 en AWS.

---

## 🔧 Tecnologías utilizadas

- **Lenguajes**: Node.js, Python, Java, Go, PHP
- **Arquitecturas**: REST, GraphQL, WebHook, SOAP, gRPC, WebSocket
- **Bases de datos**: PostgreSQL (RDS), MySQL (RDS), MongoDB Atlas, DynamoDB
- **Contenedores**: Docker + Docker Compose
- **CI/CD**: GitHub Actions
- **Despliegue**: AWS EC2 por microservicio
- **Gestión de puertos**: Uno por microservicio, sin repeticiones

---

## 🧱 Estructura por dominios

### 🔐 USUARIOS Y AUTENTICACION

| Microservicio             | Lenguaje | Arquitectura | Base de datos       | Tipo BD     | Puerto |
|---------------------------|----------|--------------|----------------------|-------------|--------|
| `auth-service`            | Node.js  | REST         | PostgreSQL (RDS)     | Relacional  | 3001   |
| `user-profile-service`    | Python   | GraphQL      | MongoDB Atlas        | NoSQL       | 3002   |
| `user-preferences-service`| PHP      | WebHook      | DynamoDB Local       | NoSQL       | 3003   |
| `role-permission-service` | Go       | REST         | PostgreSQL (RDS)     | Relacional  | 3004   |
| `address-service`         | Java     | REST         | MySQL (RDS)          | Relacional  | 3005   |

---

### 🛍️ PRODUCTOS

| Microservicio             | Lenguaje | Arquitectura | Base de datos       | Tipo BD     |
|---------------------------|----------|--------------|----------------------|-------------|
| `product-service`         | Python   | REST         | PostgreSQL           | Relacional  |
| `category-service`        | Node.js  | REST         | MySQL                | Relacional  |
| `brand-service`           | Java     | SOAP         | MySQL                | Relacional  |
| `product-variant-service` | PHP      | REST         | MongoDB              | NoSQL       |
| `product-image-service`   | Go       | gRPC         | S3 + Redis           | Cache/File  |

> *Desarrollo en curso. Se desplegarán tras finalizar `autenticacion`.*

---

### 📦 PEDIDOS Y PAGOS

| Microservicio             | Lenguaje | Arquitectura | Base de datos       | Tipo BD     |
|---------------------------|----------|--------------|----------------------|-------------|
| `order-service`           | Java     | REST         | PostgreSQL           | Relacional  |
| `order-tracking-service`  | Go       | WebSocket    | Redis                | En memoria  |
| `payment-service`         | Python   | REST         | PostgreSQL           | Relacional  |
| `refund-service`          | Node.js  | REST         | DynamoDB             | NoSQL       |
| `invoice-service`         | PHP      | WebHook      | MongoDB              | NoSQL       |

---

## 🚀 Despliegue

- Cada microservicio se encuentra en:

~/ec2/EcomVida-V1.0/<dominio>/<nombre-del-microservicio>

- Despliegue automático por GitHub Actions (`test2`)
- Secrets globales para `autenticacion`:
- `AUTH_EC2_HOST`
- `AUTH_EC2_USER`
- `AUTH_EC2_SSH_KEY`

---

## 🔁 Recomendaciones

- Usar `restart: always` en todos los `docker-compose.yml`
- Crear las tablas en RDS con `db.sql` antes del despliegue
- Mantener control de puertos sin repeticiones
- Validar logs con `docker logs <nombre>` tras deploy

---

## 📌 Estado general

| Dominio         | Estado      |
|-----------------|-------------|
| Autenticacion   | ✅ En producción (4/5 activos) |
| Productos       | 🔧 En desarrollo              |
| Pedidos y Pagos | 🔧 En desarrollo              |

---
