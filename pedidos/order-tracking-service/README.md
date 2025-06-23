# order-tracking-service

Microservicio WebSocket para publicar eventos de seguimiento de órdenes en Redis.

## 🛠️ Tecnologías

- Go 1.21
- WebSocket
- Redis
- Docker & Docker Compose
- GitHub Actions CI

## 🚀 Endpoints

- **WebSocket:** `ws://localhost:8088/ws`

Envía mensajes para ser publicados en el canal `tracking` de Redis.

## 🧪 Pruebas

Puedes probar usando Postman o con el archivo `cliente.html` en tu navegador.

## ⚙️ Variables de entorno (.env)

```env
REDIS_HOST=order-tracking-db
REDIS_PORT=6379

```

---

## Levantar Contenedor

docker compose up --build
