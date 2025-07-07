# Auth-Service

Microservicio de autenticación para el sistema Ecomvida.

## 📌 Descripción

Este microservicio permite registrar, iniciar y cerrar sesión de los usuarios. Utiliza Node.js con arquitectura RESTful y el patrón de diseño **Factory Method**. Está conectado a una base de datos PostgreSQL de tipo relacional alojada en RDS (AWS). El servicio también cuenta con su propio frontend desplegado junto al backend en una instancia EC2.

---

## 🧱 Tecnologías

| Componente      | Detalle         |
|-----------------|-----------------|
| Lenguaje        | Node.js         |
| Arquitectura    | REST            |
| Patrón de diseño| Factory Method  |
| Base de datos   | PostgreSQL (RDS)|
| Tipo de BD      | Relacional      |
| Despliegue      | EC2 (AWS)       |
| Frontend        | Sí (conectado)  |

---

## 📭 Endpoints

| Método | Endpoint     | Descripción        |
|--------|--------------|--------------------|
| POST   | `/register`  | Registro de usuario|
| POST   | `/login`     | Inicio de sesión   |
| POST   | `/logout`    | Cierre de sesión   |

---

## 🚀 Despliegue

- Backend y frontend desplegados en una instancia EC2
- Base de datos PostgreSQL gestionada a través de Amazon RDS
- Uso de Docker y GitHub actions
- Comunicación segura y eficiente entre servicios

---

## 🧑‍💻 Autor

IS Solutions - Proyecto Ecomvida  
