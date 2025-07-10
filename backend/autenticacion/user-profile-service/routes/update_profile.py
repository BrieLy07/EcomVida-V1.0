from fastapi import APIRouter, Request, Depends, HTTPException
from controllers.update_profile import actualizar_perfil
from middlewares.auth import validar_token

router = APIRouter()

@router.put("/users/me")
async def put_perfil(request: Request, usuario = Depends(validar_token)):
    data = await request.json()
    data["usuario_id"] = usuario["id"]
    return actualizar_perfil(data)

