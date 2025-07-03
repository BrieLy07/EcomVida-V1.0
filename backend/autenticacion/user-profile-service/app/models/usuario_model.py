from pydantic import BaseModel, EmailStr
from typing import Optional

class Usuario(BaseModel):
    nombre: Optional[str]
    apellido: Optional[str]
    usuario: Optional[str]
    correo: Optional[EmailStr]
    numero: Optional[str]
