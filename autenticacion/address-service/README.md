# Address Service - ECOMVIDA

Microservicio REST en Java con Spring Boot que gestiona direcciones de usuarios. Usa MySQL para almacenar datos y JWT para autenticación..

---

## 🧱 Tecnologías
- Java 17
- Spring Boot 3.5
- MySQL 8
- Docker y Docker Compose
- JWT

---

## 📦 Endpoints

| Método | Ruta              | Descripción                    |
|--------|-------------------|--------------------------------|
| GET    | `/direcciones/`   | Obtener direcciones del usuario |
| POST   | `/direcciones/`   | Crear nueva dirección           |

---

## ⚙️ Variables de entorno (.env)

```env
DB_HOST=address-db
DB_PORT=3306
DB_USER=root
DB_PASSWORD=admin123
DB_NAME=addressdb
