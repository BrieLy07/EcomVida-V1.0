# 🚚 Inventario y Logística - Ecomvida

Este dominio gestiona el control de inventarios, almacenes, envíos, entregas y devoluciones en Ecomvida. Está compuesto por 5 microservicios independientes, cada uno desarrollado con su propio lenguaje, arquitectura y patrón de diseño.

---

## 🧩 Microservicios

| Microservicio         | Lenguaje | Arquitectura | Patrón de Diseño       | Base de Datos        | Tipo     | Puerto |
|-----------------------|----------|--------------|------------------------|-----------------------|----------|--------|
| inventory-service     | Python   | REST         | Repository             | MySQL                 | Relacional | 3021   |
| warehouse-service     | Go       | REST         | Service Locator        | PostgreSQL            | Relacional | 3022   |
| shipment-service      | Node.js  | WebHook      | Mediator               | Elasticsearch         | Otros     | 3023   |
| delivery-service      | Node.js  | REST         | State                  | MongoDB Atlas         | NoSQL     | 3024   |
| return-service        | Python   | REST         | Chain of Responsibility| PostgreSQL            | Relacional | 3020   |

---

## 📦 Funcionalidades

- Gestión de inventarios (altas, bajas, actualizaciones).
- Administración de almacenes físicos.
- Gestión de envíos y seguimiento en tiempo real.
- Registro y control de entregas.
- Procesamiento de devoluciones con validaciones por cadena de responsabilidad.

---

## 🛠️ Detalles Técnicos

- Cada microservicio está en su propia carpeta con su `README.md`.
- Se aplican principios SOLID, KISS y DRY.
- Todos los endpoints están separados por archivo.
- Cada microservicio se dockeriza individualmente para despliegue en EC2.
- Se usan `.env` para configuración y secretos en GitHub.
- MongoDB Atlas y Redis son usados en la nube; Elasticsearch y Dynamo en local.

---

## 🚀 Despliegue

Cada microservicio será desplegado de forma independiente usando:

- Docker y GitHub Actions.
- EC2 + API Gateway + Load Balancer.
- Todo optimizado para el Free Tier de AWS.