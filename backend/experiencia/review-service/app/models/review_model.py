from typing import Optional
from pydantic import BaseModel, Field
from datetime import datetime

class Review(BaseModel):
    id: Optional[str] = Field(alias="_id")
    producto_id: str
    usuario_id: str
    titulo: str
    contenido: str
    rating: int
    fecha_creacion: Optional[datetime] = None
