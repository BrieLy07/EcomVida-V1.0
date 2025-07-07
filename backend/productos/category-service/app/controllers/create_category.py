from fastapi import APIRouter
from app.models.categoria import Categoria
from app.repositories.category_repository import CategoriaRepository

router = APIRouter()

@router.post("/categories")
def crear_categoria(data: dict):
    nueva = Categoria(None, data["nombre"], data["descripcion"])
    creada = CategoriaRepository.crear(nueva)
    return creada.to_dict()
