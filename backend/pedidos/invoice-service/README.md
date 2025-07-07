# 🧾 Invoice Service - Ecomvida

Microservicio para generar y gestionar facturas de pedidos. Implementado en Node.js, arquitectura REST, usando **MongoDB Atlas** como base de datos y el patrón de diseño **Strategy**.

---

## 🧩 Descripción del Servicio

Permite emitir facturas simples o detalladas para pedidos, obtenerlas por ID u orden, actualizarlas o eliminarlas.

---

## 🏗️ Estructura del Proyecto

invoice-service/
├── app/
│ ├── controllers/ # Controladores (uno por endpoint)
│ ├── database/ # Conexión MongoDB
│ ├── models/ # Modelo Factura
│ └── strategies/ # Estrategias de generación
├── main.js # Inicializa Express
├── Dockerfile # Imagen optimizada
├── .env.example # Variables de entorno
├── package.json # Dependencias
└── README.md


---

## 🔁 Endpoints

| Método | Endpoint                     | Descripción                          |
|--------|------------------------------|--------------------------------------|
| POST   | `/invoices`                  | Generar nueva factura                |
| GET    | `/invoices/:id`              | Obtener factura por ID               |
| GET    | `/invoices/order/:orderId`   | Obtener factura de un pedido         |
| PUT    | `/invoices/:id`              | Actualizar factura                   |
| DELETE | `/invoices/:id`              | Eliminar factura                     |

---

## 🧠 Patrón de Diseño

**Strategy**  
Permite generar distintos tipos de factura mediante clases como `FacturaSimple`, `FacturaDetallada`.

---

## ⚙️ Variables de Entorno

```env
PORT=3015
MONGO_URI=mongodb+srv://<usuario>:<password>@<cluster>.mongodb.net/ecomvida
```

## 📦 Dependencias

```express
mongoose
dotenv
```
---

## 🚀 Despliegue
El microservicio está listo para conectarse a MongoDB Atlas directamente. También puede adaptarse a MongoDB local si se requiere.

---

### 📄 `.gitignore`

```gitignore
node_modules/
.env
*.log
.DS_Store
