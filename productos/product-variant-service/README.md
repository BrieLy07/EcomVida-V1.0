# Product Variant Service - ECOMVIDA

Microservicio REST en PHP (Slim Framework) con MongoDB para gestionar variantes de productos como componentes de PC y periféricos.

---

## 📦 Endpoints

| Método | Ruta                  | Descripción                         |
|--------|------------------------|-------------------------------------|
| POST   | /variantes             | Crear una nueva variante            |
| GET    | /variantes             | Listar todas las variantes          |
| GET    | /variantes/{id}        | Obtener una variante por su ID      |

---

## 🔧 Ejemplo JSON para crear variante

```json
{
  "producto_id": "gpu-4090",
  "color": "negro",
  "memoria": "24GB GDDR6X",
  "stock": 5,
  "precio": 1899.99
}

```

---

## 🐳 Docker

docker compose up --build
