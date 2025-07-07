# 🏢 Warehouse Service - Ecomvida

Microservicio responsable de la gestión de almacenes físicos en Ecomvida. Desarrollado en **Go** usando **Gin** con arquitectura REST y patrón de diseño **Service Locator**. Utiliza base de datos **PostgreSQL**.

---

## 🧩 Descripción del Servicio

Permite registrar, consultar, actualizar y eliminar almacenes. Toda la lógica está desacoplada en capas para facilitar mantenimiento y pruebas.

---

## 🏗️ Estructura del Proyecto

warehouse-service/
├── cmd/
│ └── main.go # Punto de entrada
├── config/ # Conexión a PostgreSQL
├── controller/ # Endpoints separados
├── locator/ # Service Locator
├── model/ # Estructura del almacén
├── repository/ # Lógica de acceso a datos
├── go.mod / go.sum # Dependencias
├── Dockerfile
├── .env.example
└── README.md

---

## 🔁 Endpoints

| Método | Endpoint             | Descripción                      |
|--------|----------------------|----------------------------------|
| GET    | `/warehouses`        | Listar todos los almacenes       |
| GET    | `/warehouses/:id`    | Obtener detalles de un almacén   |
| POST   | `/warehouses`        | Crear un nuevo almacén           |
| PUT    | `/warehouses/:id`    | Actualizar información           |
| DELETE | `/warehouses/:id`    | Eliminar un almacén              |

---

## 🧠 Patrón de Diseño

**Service Locator**  
Se centraliza la gestión de dependencias y servicios, facilitando su reutilización sin acoplamiento directo.

---

## ⚙️ Variables de Entorno

```env
PORT=3017
POSTGRES_URL=postgres://usuario:password@host:puerto/nombre_basedatos?sslmode=disable
```

## 🚀 Despliegue
Este microservicio está preparado para ejecutarse en una instancia EC2 con PostgreSQL en RDS. Dockeriza solo lo necesario para producción.


