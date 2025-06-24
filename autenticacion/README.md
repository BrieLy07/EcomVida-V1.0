# 🔐 Dominio: USUARIOS Y AUTENTICACION

Este dominio agrupa todos los microservicios relacionados con autenticación, perfil de usuario, gestión de roles, direcciones y preferencias. Cada microservicio se despliega individualmente sobre EC2 y usa su propia base de datos.

---

## 🧱 Microservicios del dominio

| Microservicio             | Lenguaje | Arquitectura | Base de datos       | Tipo BD     | Puerto |
|---------------------------|----------|--------------|----------------------|-------------|--------|
| `auth-service`            | Node.js  | REST         | PostgreSQL (RDS)     | Relacional  | 3001   |
| `user-profile-service`    | Python   | GraphQL      | MongoDB Atlas        | NoSQL       | 3002   |
| `user-preferences-service`| PHP      | WebHook      | DynamoDB Local       | NoSQL       | 3003   |
| `role-permission-service` | Go       | REST         | PostgreSQL (RDS)     | Relacional  | 3004   |
| `address-service`         | Java     | REST         | MySQL (RDS)          | Relacional  | 3005   |

---

## 🚀 Despliegue

Cada microservicio está ubicado en EC2 bajo:

~/ec2/EcomVida-V1.0/autenticacion/<microservicio>


- Todos los servicios se levantan con Docker Compose
- Usan puertos no repetidos (3001 a 3005)
- Configurados con `restart: always` para reinicio automático
- Base de datos externa en RDS (excepto preferencias que es DynamoDB local)

---

## ⚙ GitHub Actions

Despliegue automático con cada push a la rama `test2`.

Secrets compartidos en todo el dominio:
- `AUTH_EC2_HOST`
- `AUTH_EC2_USER`
- `AUTH_EC2_SSH_KEY`

---

## 📌 Estado de los microservicios

- ✅ `auth-service`: desplegado y probado
- ✅ `user-profile-service`: desplegado y funcional
- ✅ `role-permission-service`: operativo con RDS
- ✅ `address-service`: adaptado a producción
- 🕓 `user-preferences-service`: pendiente por conectar con DynamoDB local

---

## 📦 Consideraciones

- Las tablas necesarias se crean en RDS con archivos `db.sql`
- Los contenedores se levantan con:
```bash
docker-compose down
docker-compose up -d --build
