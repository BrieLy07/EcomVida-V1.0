# Product Image Service - ECOMVIDA

Microservicio REST en Go para subir y obtener imágenes de productos, usando MinIO (S3) y Redis.

---

## 📦 Endpoints

| Método | Ruta                   | Descripción                       |
|--------|-------------------------|-----------------------------------|
| POST   | /imagenes               | Subir una imagen                  |
| GET    | /imagenes/{nombre}      | Obtener URL firmada de la imagen |

---

## 🔧 Subida de imagen

Enviar una imagen por `form-data` con el campo `imagen`.

---

## 📂 Variables de entorno (.env)

```env
MINIO_ENDPOINT=minio:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin
MINIO_BUCKET=imagenes-producto

REDIS_ADDR=redis:6379

```

---

## 🐳 Docker

docker compose up --build

