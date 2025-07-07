# 🔄 Return Service - Ecomvida

Microservicio responsable de gestionar las devoluciones de productos en Ecomvida. Desarrollado en **Python**, arquitectura **REST**, patrón de diseño **Chain of Responsibility** y persistencia en **PostgreSQL**.

---

## 🧩 Descripción del Servicio

Permite registrar, consultar, actualizar o eliminar devoluciones. Utiliza una cadena de validaciones antes de ejecutar la lógica principal.

---

## 🏗️ Estructura del Proyecto

return-service/
├── app/
│ ├── controllers/ # Endpoints separados
│ ├── database/ # Conexión a PostgreSQL
│ ├── handlers/ # Chain of Responsibility
│ └── models/ # Modelo SQLAlchemy
├── .env.example
├── Dockerfile
├── requirements.txt
├── main.py
└── README.md


---

## 🔁 Endpoints

| Método | Endpoint         | Descripción                        |
|--------|------------------|------------------------------------|
| GET    | /returns         | Listar todas las devoluciones      |
| GET    | /returns/:id     | Obtener detalles de una devolución |
| POST   | /returns         | Registrar una nueva devolución     |
| PUT    | /returns/:id     | Actualizar estado de devolución    |
| DELETE | /returns/:id     | Eliminar una devolución            |

---

## 🧠 Patrón de Diseño

**Chain of Responsibility**  
Cada validación (como motivo y estado) se realiza en una cadena modular y extensible.

---

## ⚙️ Variables de Entorno

```env
PORT=3020
POSTGRES_URL=postgresql://usuario:contraseña@host:puerto/ecomvida
```

## 🚀 Despliegue
Este microservicio está listo para su ejecución en contenedor Docker o EC2. Incluye migración automática de la tabla returns mediante SQLAlchemy.