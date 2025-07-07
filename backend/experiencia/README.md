# Dominio: Experiencia del Usuario

Este dominio contiene microservicios enfocados en mejorar la experiencia directa del cliente en el eCommerce Ecomvida, brindando funcionalidades como reseñas, listas de deseos, carrito, búsquedas inteligentes y recomendaciones personalizadas.

---

## 🧩 Microservicios del Dominio

| Microservicio             | Lenguaje | Arquitectura | Patrón de Diseño       | Base de Datos     | Tipo      |
|---------------------------|----------|--------------|------------------------|-------------------|-----------|
| review-service            | Python   | GraphQL      | Decorator              | MongoDB Atlas     | NoSQL     |
| wishlist-service          | Node.js  | REST         | Observer               | DynamoDB Local    | NoSQL     |
| cart-service              | Go       | REST         | Memento                | Redis             | En memoria|
| search-service            | Rust     | REST         | Cache Proxy            | Elasticsearch     | Otros     |
| recommendation-service    | Python   | gRPC         | Chain of Responsibility| PostgreSQL (RDS)  | Relacional|

---

## ✅ Estado de los Servicios

- Todos los microservicios han sido desarrollados con su backend completo.
- Cada endpoint está en su propio archivo para facilitar mantenimiento.
- Se ha usado `.env` y `secrets` para variables sensibles.
- Los servicios que no pueden estar aún en la nube (por condiciones de laboratorio) se ejecutan de forma local (DynamoDB local, Redis local).

---

## 🚀 Despliegue

- MongoDB Atlas se usa como base de datos NoSQL para `review-service`.
- Redis está en entorno local y será evaluado para su despliegue en EC2 si aplica.
- PostgreSQL en RDS usado en `recommendation-service`.
- Cada microservicio está listo para ser empaquetado y desplegado en EC2.
- Se utilizarán GitHub Actions y configuración por `.yml` para automatización.

---

## 🧪 Pruebas

- Cada microservicio ha sido probado individualmente con Postman (REST) o pruebas manuales/gRPCui (gRPC/GraphQL).
- Las respuestas y funcionamiento están validados conforme a sus requerimientos.

---

## 📝 Notas

- Todos los servicios cumplen con los patrones de diseño asignados y siguen las buenas prácticas SOLID/KISS.
- El despliegue se realizará tras completar todos los dominios del proyecto Ecomvida.
