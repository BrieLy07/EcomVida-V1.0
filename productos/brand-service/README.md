# Brand Service - ECOMVIDA

Microservicio SOAP en Java (Spring Boot + Apache CXF) para gestión de marcas de productos.

---

## 📦 Métodos SOAP

| Método         | Parámetro       | Descripción                  |
|----------------|-----------------|------------------------------|
| crearMarca     | nombre: String   | Crea una nueva marca         |
| listarMarcas   | —               | Lista todas las marcas       |
| obtenerMarca   | id: Long        | Devuelve una marca por ID    |

---

## 🧪 URL WSDL

http://localhost:8083/ws/MarcaEndpoint?wsdl

---

## ⚙️ Variables de entorno (.env)

```env
DB_HOST=brand-db
DB_PORT=3306
DB_NAME=branddb
DB_USER=root
DB_PASSWORD=admin123

```

---

## 🐳 Docker

docker compose up --build
