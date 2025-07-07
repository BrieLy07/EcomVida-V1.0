from fastapi import APIRouter, HTTPException
from app.controllers.actualizar_usuario_controller import actualizar_usuario
from app.models.usuario_model import Usuario

router = APIRouter()

@router.put("/users/{id}")
def actualizar(id: str, datos: Usuario):
    exito = actualizar_usuario(id, datos.dict(exclude_unset=True))
    if not exito:
        raise HTTPException(status_code=404, detail="Usuario no encontrado o sin cambios")
    return {"mensaje": "Usuario actualizado correctamente"}
