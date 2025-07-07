# 📦 Delivery Service - Ecomvida

Microservicio encargado de gestionar las entregas de productos en la plataforma Ecomvida. Desarrollado en **Node.js**, arquitectura **REST**, patrón de diseño **State** y persistencia en **MongoDB Atlas**.

---

## 🧩 Descripción del Servicio

Permite registrar, consultar, actualizar y eliminar entregas. Implementa lógica de cambio de estado utilizando el patrón **State**.

---

## 🏗️ Estructura del Proyecto

delivery-service/
├── src/
│ ├── controllers/ # Endpoints separados
│ ├── database/ # Conexión a MongoDB Atlas
│ ├── models/ # Modelo Mongoose
│ └── states/ # Implementación del patrón State
├── .env.example
├── Dockerfile
├── package.json
├── .gitignore
└── README.md

---

## 🔁 Endpoints

| Método | Endpoint            | Descripción                         |
|--------|---------------------|-------------------------------------|
| GET    | `/deliveries`       | Listar entregas realizadas/pendientes |
| GET    | `/deliveries/:id`   | Obtener detalles de una entrega     |
| POST   | `/deliveries`       | Registrar una nueva entrega         |
| PUT    | `/deliveries/:id`   | Actualizar estado de entrega        |
| DELETE | `/deliveries/:id`   | Eliminar una entrega                |

---

## 🧠 Patrón de Diseño

**State**  
Cada estado de la entrega (`pendiente`, `en_transito`, `entregado`, `cancelado`) define su propia lógica para transiciones válidas.

---

## ⚙️ Variables de Entorno

```env
PORT=3019
MONGO_URI=mongodb+srv://<usuario>:<contraseña>@cluster.mongodb.net/ecomvida

```

## Despliegue
Este microservicio está dockerizado y listo para ejecutarse en EC2.

MongoDB Atlas se utiliza como base de datos NoSQL en la nube.

