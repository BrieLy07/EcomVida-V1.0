from fastapi import APIRouter, HTTPException
from app.controllers.obtener_usuario_controller import obtener_usuario

router = APIRouter()

@router.get("/users/{id}")
def obtener(id: str):
    usuario = obtener_usuario(id)
    if not usuario:
        raise HTTPException(status_code=404, detail="Usuario no encontrado")
    return usuario
