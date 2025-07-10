from fastapi import APIRouter, Request, HTTPException
from database.conexion import ConexionMongo

def obtener_perfil(usuario_id: str):
    db = ConexionMongo.obtener_conexion()
    perfil = db.perfiles.find_one({"usuario_id": usuario_id}, {"_id": 0})
    if not perfil:
        raise HTTPException(status_code=404, detail="Perfil no encontrado")
    return perfil
