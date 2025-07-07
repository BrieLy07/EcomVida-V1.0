# 🚚 Shipment Service - Ecomvida

Microservicio responsable de gestionar los envíos realizados en la plataforma Ecomvida. Desarrollado en **Node.js**, utiliza arquitectura **WebHook**, patrón de diseño **Mediator**, y almacena los datos en **Elasticsearch**.

---

## 🧩 Descripción del Servicio

Permite registrar, consultar, actualizar y eliminar información de los envíos. Toda la lógica se comunica a través de eventos mediante un Mediador que centraliza las acciones.

---

## 🏗️ Estructura del Proyecto

shipment-service/
├── src/
│ ├── controllers/ # Endpoints separados
│ ├── mediator/ # Mediador de eventos
│ ├── services/ # Lógica por evento
│ └── utils/ # Conexión a Elasticsearch
├── .env.example
├── Dockerfile
├── package.json
├── .gitignore
└── README.md

---

## 🔁 Endpoints

| Método | Endpoint            | Descripción                          |
|--------|---------------------|--------------------------------------|
| GET    | `/shipments`        | Listar todos los envíos              |
| GET    | `/shipments/:id`    | Obtener detalles de un envío         |
| POST   | `/shipments`        | Crear un nuevo envío                 |
| PUT    | `/shipments/:id`    | Actualizar estado o destino del envío|
| DELETE | `/shipments/:id`    | Cancelar o eliminar un envío         |

---

## 🧠 Patrón de Diseño

**Mediator**  
Centraliza la comunicación entre controladores y servicios a través de eventos, desacoplando la lógica de negocio.

---

## ⚙️ Variables de Entorno

```env
PORT=3018
ELASTIC_NODE=http://localhost:9200
```

## 🚀 Despliegue
Este microservicio está dockerizado y listo para ejecutarse en EC2.
Elasticsearch se ejecuta localmente en contenedor durante el desarrollo.

