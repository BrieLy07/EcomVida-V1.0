# 🏠 Address Service – EcomVida

Microservicio RESTful en Go que gestiona las **direcciones** de los usuarios registrados. Pertenece al dominio **Usuarios y Autenticación** del sistema EcomVida.

---

## 🚀 Tecnología

| Componente     | Valor              |
|----------------|--------------------|
| Lenguaje       | Go (Golang)        |
| Arquitectura   | REST               |
| Patrón diseño  | Repository         |
| Base de datos  | MySQL (RDS)        |
| Puerto         | 3003               |
| Autenticación  | JWT (middleware)   |

---

## 📦 Endpoints

| Método | Endpoint                        | Descripción                     |
|--------|----------------------------------|----------------------------------|
| GET    | `/users/:id/addresses`          | Listar direcciones del usuario  |
| POST   | `/users/:id/addresses`          | Crear nueva dirección           |
| PUT    | `/addresses/:addressId`         | Actualizar dirección existente  |
| DELETE | `/addresses/:addressId`         | Eliminar dirección              |

> Todos los endpoints requieren **token JWT** en el header:  
> `Authorization: Bearer <token>`

---

## ⚙️ Variables de entorno `.env`

```env
PORT=3003

DB_HOST=address-db.cg9go1thnm7i.us-east-1.rds.amazonaws.com
DB_PORT=3306
DB_USER=admin
DB_PASSWORD=tu_contraseña
DB_NAME=address_db

JWT_SECRET=clave_secreta_address

```

## 🛠️ Inicializar base de datos

Ejecutar el siguiente script en tu instancia RDS MySQL:

CREATE DATABASE IF NOT EXISTS address_db;

USE address_db;

CREATE TABLE IF NOT EXISTS direcciones (
    id INT AUTO_INCREMENT PRIMARY KEY,
    usuario_id VARCHAR(255) NOT NULL,
    direccion TEXT NOT NULL,
    ciudad VARCHAR(100),
    provincia VARCHAR(100),
    pais VARCHAR(100),
    codigo_postal VARCHAR(20),
    telefono VARCHAR(20)
);

---

## Docker

docker build -t address-service .

docker run -d --name address-service -p 3003:3003 --env-file .env address-service

---

## Integración

Se comunica con los siguientes microservicios:

Microservicio	Propósito
auth-service	Validación del token JWT
user-profile-service	Lectura de datos de usuario

