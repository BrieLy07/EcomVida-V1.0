# Dominio: Autenticación – EcomVida

Este dominio agrupa los microservicios responsables de la autenticación y gestión de usuarios dentro del ecosistema EcomVida. Cada servicio está aislado, optimizado para despliegue en EC2 con Docker, y utiliza diferentes motores de base de datos según su función.

---

## 📦 Microservicios del dominio

| Microservicio             | Lenguaje | Puerto | Base de Datos | Descripción |
|--------------------------|----------|--------|----------------|-------------|
| auth-service             | Node.js  | 3001   | PostgreSQL     | Registro, login y validación de credenciales. |
| user-profile-service     | Python   | 3002   | MongoDB Atlas  | Gestión del perfil de usuario. |
| address-service          | Go       | 3003   | MySQL (RDS)    | Gestión de direcciones de usuario. |

---

## 🚀 Estructura del Proyecto

backend/
└── autenticacion/
├── auth-service/
├── user-profile-service/
└── address-service/

Cada microservicio incluye:

- Código limpio y modular
- Dockerfile optimizado para producción
- Script SQL o conexión a base NoSQL
- Variables de entorno gestionadas con `.env`
- Workflows en GitHub Actions para CI/CD

---

## ⚙️ Despliegue en EC2

Cada microservicio se despliega en su propia instancia EC2 mediante GitHub Actions. El flujo es:

1. Commit a rama `test2.0`
2. GitHub Actions construye la imagen Docker
3. Se despliega automáticamente en la instancia correspondiente
4. Las bases de datos están en RDS (PostgreSQL y MySQL) o MongoDB Atlas

---

## 📄 Workflows disponibles

- `auth-service`: `deploy-auth-service.yml`
- `user-profile-service`: `deploy-user-profile.yml`
- `address-service`: `deploy-address-service.yml`

Cada workflow:
- Ejecuta tests unitarios
- Construye imagen Docker (`brielys/...`)
- Despliega vía SSH usando claves privadas seguras
- Asigna puertos únicos

---

## 🔐 Seguridad

- JWT aplicado en `auth-service` para autenticación
- Comunicación entre microservicios se realiza por rutas privadas
- MongoDB y RDS protegidos por credenciales y sin acceso público
- `.env` excluido del control de versiones

---

## 📬 Contacto

Proyecto desarrollado por el equipo de **EcomVida**.  
Para dudas técnicas, escribir a: [soporte@ecomvida.ec](mailto:soporte@ecomvida.ec)