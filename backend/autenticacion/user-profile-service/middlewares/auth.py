from fastapi import Request, HTTPException
from jose import jwt, JWTError
import os

SECRET_KEY = os.getenv("JWT_SECRET", "miclaveultrasecreta")

async def validar_token(request: Request):
    auth = request.headers.get("Authorization")
    if not auth or not auth.startswith("Bearer "):
        raise HTTPException(status_code=401, detail="Token requerido")

    token = auth.split(" ")[1]
    try:
        datos = jwt.decode(token, SECRET_KEY, algorithms=["HS256"])
        return datos
    except JWTError:
        raise HTTPException(status_code=403, detail="Token inválido")
