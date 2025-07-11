# Frontend – Autenticación | EcomVida

Este frontend corresponde al módulo de autenticación del sistema EcomVida. Permite a los usuarios registrarse, iniciar sesión, visualizar y editar su perfil, así como gestionar su dirección.

---

## 🧩 Funcionalidades implementadas

- **Pantalla principal con navegación**
- **Registro de usuarios** con campos:
  - nombre, apellido, usuario, correo, teléfono, contraseña, rol
- **Inicio de sesión**
- **Vista de perfil** (solo lectura)
- **Editar perfil y dirección**
- **Eliminar dirección**
- Toda la interfaz mantiene un **estilo visual unificado simple y claro**

---

## 🚀 Scripts de desarrollo

```bash
npm install      # Instala dependencias
npm start        # Corre el frontend en localhost:3000
npm run build    # Compila a producción
```

## 🐳 Docker
El frontend está preparado para despliegue con Docker.

Dockerfile incluido en auth-front/:

docker build -t auth-frontend .
docker run -d -p 80:80 auth-frontend

## 🌐 Despliegue en producción
El frontend se construye en modo producción (npm run build) y se sirve con Nginx en una instancia EC2, usando un workflow de GitHub Actions que:

Elimina contenedores antiguos

Actualiza el repositorio

Construye imagen Docker

Expone en puerto 80

## 📄 Workflows relacionados
Despliegue EcomVida Auth Services: despliega auth-service y auth-front juntos en EC2

Usa secretos: AUTH_EC2_HOST, AUTH_EC2_USER, AUTH_EC2_SSH_KEY