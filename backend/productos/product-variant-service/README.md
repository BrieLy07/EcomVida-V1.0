# 🎨 Product Variant Service - Ecomvida

Microservicio encargado de gestionar las variantes de productos en Ecomvida, como color, talla, disponibilidad y precio adicional.

---

## 🧩 Descripción del Servicio

Permite crear, listar, actualizar y eliminar variantes asociadas a productos específicos. Desarrollado en **Node.js**, usa **MongoDB** y aplica el patrón de diseño **Builder** para construir variantes de forma flexible.

---

## 🏗️ Estructura del Proyecto

product-variant-service/
├── src/
│ ├── controllers/ # Endpoints separados
│ ├── models/ # Esquema Variante
│ ├── builders/ # Patrón Builder para construir variantes
│ ├── routes/ # Rutas agrupadas
│ └── database.js # Conexión a MongoDB
├── server.js
├── package.json
├── Dockerfile
├── .env.example
└── README.md


---

## 🧪 Endpoints Disponibles

| Método | Endpoint                                 | Descripción                          |
|--------|------------------------------------------|--------------------------------------|
| GET    | `/variants`                              | Listar todas las variantes           |
| GET    | `/variants/:id`                          | Obtener una variante por ID          |
| POST   | `/products/:productId/variants`          | Crear una variante para un producto |
| PUT    | `/variants/:id`                          | Actualizar una variante              |
| DELETE | `/variants/:id`                          | Eliminar una variante                |

---

## 🧱 Patrón de Diseño Utilizado

**Builder:**  
Permite construir objetos `Variante` de forma flexible y clara antes de ser guardados en la base de datos.

---

## ⚙️ Variables de Entorno

Ver `.env.example`:

PORT=3009
MONGODB_URI=mongodb://localhost:27017/variants

---

## 🐳 Dockerización

El microservicio corre en el puerto `3009` y está optimizado para producción con `Dockerfile`. Listo para desplegar en EC2 o entorno local.

---

## ✅ Pruebas

Todos los endpoints han sido probados exitosamente usando Postman.