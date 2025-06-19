import os
import jwt
from flask import request
from dotenv import load_dotenv

load_dotenv()
SECRET = os.getenv("JWT_SECRET")

def obtener_usuario_autenticado():
    auth_header = request.headers.get("Authorization")
    if not auth_header or not auth_header.startswith("Bearer "):
        return None

    token = auth_header.split(" ")[1]
    try:
        payload = jwt.decode(token, SECRET, algorithms=["HS256"])
        return payload.get("id")  # ID del usuario
    except jwt.ExpiredSignatureError:
        return None
    except jwt.InvalidTokenError:
        return None
