from fastapi import APIRouter, HTTPException, Depends
from controllers.get_profile import obtener_perfil
from middlewares.auth import validar_token

router = APIRouter()

@router.get("/users/me")
def get_perfil(usuario = Depends(validar_token)):
    return obtener_perfil(usuario["id"])

