# 📦 Inventory Service - Ecomvida

Microservicio encargado de la gestión del inventario de productos en Ecomvida. Implementado en **Python** con **FastAPI**, arquitectura REST y el patrón de diseño **Repository**.

---

## 🧩 Descripción del Servicio

Permite listar, registrar, actualizar y eliminar registros de inventario. Se conecta a una base de datos **MySQL** y crea automáticamente la tabla al iniciar.

---

## 🏗️ Estructura del Proyecto

inventory-service/
├── app/
│ ├── controllers/ # Controladores por endpoint
│ ├── database/ # Conexión MySQL
│ ├── models/ # Modelo Inventario
│ └── repository/ # Lógica con patrón Repository
├── main.py # Inicializador FastAPI
├── Dockerfile # Imagen optimizada
├── .env.example # Variables de entorno
├── requirements.txt # Dependencias
└── README.md


---

## 🔁 Endpoints

| Método | Endpoint          | Descripción                          |
|--------|-------------------|--------------------------------------|
| GET    | `/inventory`      | Listar todos los productos           |
| GET    | `/inventory/:id`  | Obtener inventario por ID            |
| POST   | `/inventory`      | Registrar nuevo inventario           |
| PUT    | `/inventory/:id`  | Actualizar cantidad o estado         |
| DELETE | `/inventory/:id`  | Eliminar registro de inventario      |

---

## 🧠 Patrón de Diseño

**Repository**  
Se separa la lógica de acceso a datos en una clase independiente (`InventarioRepository`) para facilitar mantenimiento y pruebas.

---

## ⚙️ Variables de Entorno

```env
PORT=3016
MYSQL_URL=mysql+pymysql://usuario:password@host:puerto/nombre_basedatos
```

## 🚀 Despliegue
Este servicio está preparado para ser ejecutado en una instancia EC2 conectada a una base de datos MySQL en RDS. Solo se empaquetan archivos necesarios para producción.

