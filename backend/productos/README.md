# 🧱 Dominio Productos - Ecomvida

Este dominio agrupa todos los microservicios encargados de gestionar la información relacionada a productos dentro del ecosistema Ecomvida.

---

## 🧩 Microservicios

| Microservicio              | Lenguaje | Arquitectura | Patrón Diseño | Base de Datos | Tipo     |
|----------------------------|----------|--------------|----------------|---------------|----------|
| product-service            | Python   | REST         | Repository     | PostgreSQL    | Relacional |
| category-service           | Python   | REST         | Composite      | PostgreSQL    | Relacional |
| brand-service              | Python   | REST         | Adapter        | PostgreSQL    | Relacional |
| product-variant-service    | Node.js  | REST         | Builder        | MongoDB       | NoSQL      |
| product-image-service      | Go       | gRPC         | Proxy          | MinIO + Redis | S3 + Cache |

---

## 🔁 Flujo de Trabajo

Cada microservicio:

- Tiene su propia carpeta.
- Expone sus endpoints/documentación en su propio `README.md`.
- Sigue un patrón de diseño distinto y optimizado.
- Está dockerizado individualmente para producción.
- Utiliza su propio puerto sin repetir.
- Se comunica con su base de datos asignada.
- Expone endpoints REST o gRPC según corresponda.

---

## 📦 Estructura

productos/
├── product-service/
├── category-service/
├── brand-service/
├── product-variant-service/
├── product-image-service/
└── README.md (este archivo)

---

## 🧪 Pruebas

Todos los endpoints han sido probados con Postman (REST) y BloomRPC (gRPC). Listos para despliegue en AWS.

