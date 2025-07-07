from fastapi import APIRouter
from app.models.marca import Marca
from app.repositories.marca_repository import MarcaRepository

router = APIRouter()

@router.post("/brands")
def crear_marca(data: dict):
    nueva = Marca(None, data["nombre"], data["descripcion"])
    creada = MarcaRepository.crear(nueva)
    return creada.to_dict()
