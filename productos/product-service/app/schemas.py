from pydantic import BaseModel

class ProductoBase(BaseModel):
    nombre: str
    descripcion: str | None = None
    precio: float

class ProductoCrear(ProductoBase):
    pass

class ProductoRespuesta(ProductoBase):
    id: int

    class Config:
        orm_mode = True
