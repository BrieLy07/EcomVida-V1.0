from fastapi import HTTPException
from database.conexion import ConexionMongo
from models.perfil import crear_objeto_perfil

def actualizar_perfil(data: dict):
    if not data.get("usuario_id"):
        raise HTTPException(status_code=400, detail="Falta usuario_id")
    
    db = ConexionMongo.obtener_conexion()
    perfil = crear_objeto_perfil(data)

    resultado = db.perfiles.update_one(
        {"usuario_id": perfil["usuario_id"]},
        {"$set": perfil},
        upsert=True
    )

    if resultado.matched_count == 0 and resultado.upserted_id is None:
        raise HTTPException(status_code=500, detail="No se pudo actualizar el perfil")

    return {"mensaje": "Perfil actualizado", "perfil": perfil}
