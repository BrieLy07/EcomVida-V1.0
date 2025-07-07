# 🖼️ Product Image Service - Ecomvida

Microservicio gRPC para gestión de imágenes de productos. Usa almacenamiento tipo S3 (MinIO en local) y caché en Redis. Aplica el patrón de diseño **Proxy**.

---

## 🧩 Descripción del Servicio

Este servicio permite subir, obtener, listar y eliminar imágenes asociadas a productos. Usa un proxy que controla el acceso al almacenamiento (MinIO) y a la caché (Redis) para mejorar el rendimiento.

---

## 🏗️ Estructura del Proyecto

product-image-service/
├── app/
│ ├── cache/ # Operaciones con Redis (caché de metadatos)
│ ├── handlers/ # Implementación del servicio gRPC
│ ├── proxy/ # Proxy hacia MinIO + Redis
│ ├── s3/ # Conexión y operaciones con MinIO
│ └── utils/ # Funciones auxiliares
├── proto/
│ └── image.proto # Definición del servicio gRPC
│ └── imagepb/ # Código generado por protoc
├── server.go
├── Dockerfile
├── go.mod
├── .env.example
└── README.md

---

## 🔁 Endpoints gRPC

| Método RPC              | Descripción                                   |
|-------------------------|-----------------------------------------------|
| UploadImage             | Subir una imagen                              |
| GetImage                | Obtener imagen por ID                         |
| GetImagesByProduct      | Listar imágenes de un producto                |
| DeleteImage             | Eliminar una imagen                           |

---

## 🎯 Patrón de Diseño Utilizado

**Proxy**  
El servicio actúa como intermediario entre el cliente y MinIO/Redis, controlando las solicitudes, optimizando el acceso y ocultando la complejidad del backend.

---

## ⚙️ Variables de Entorno

Ver archivo `.env.example`:

PORT=30010

MinIO (emulando S3)
S3_ENDPOINT=http://minio:9000
S3_ACCESS_KEY=minioadmin
S3_SECRET_KEY=minioadmin
S3_BUCKET=imagenes-productos

Redis
REDIS_HOST=redis
REDIS_PORT=6379


---

## 🐳 Dockerización

Este microservicio se ejecuta como contenedor gRPC. Puedes integrarlo con contenedores de MinIO y Redis en un mismo `docker-compose`.

---

## ✅ Pruebas

Para probarlo se puede usar un cliente gRPC (como BloomRPC o Postman con plugin gRPC), apuntando al puerto `30010`.

---


