def crear_objeto_perfil(data):
    return {
        "usuario_id": data.get("usuario_id"),
        "nombre": data.get("nombre"),
        "apellido": data.get("apellido"),
        "correo": data.get("correo"),
        "numero": data.get("numero")
    }
