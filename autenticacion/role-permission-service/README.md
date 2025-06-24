# role-permission-service

Microservicio del dominio `autenticacion` encargado de la gestión de roles y permisos. Desarrollado en Go con arquitectura REST y base de datos PostgreSQL en RDS.

---

### 🧪 Endpoints

**Roles**
- `POST /roles` → Crear rol
- `GET /roles` → Listar roles

**Permisos**
- `POST /permisos` → Crear permiso
- `GET /permisos` → Listar permisos

---

### ⚙ Variables de entorno

```env
PORT=3004
DB_HOST=auth-db.cg9go1thnm7i.us-east-1.rds.amazonaws.com
DB_PORT=5432
DB_NAME=authdb
DB_USER=postgres
DB_PASSWORD=admin123
JWT_SECRET=miclavesecreta
