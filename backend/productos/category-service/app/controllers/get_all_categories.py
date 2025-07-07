from fastapi import APIRouter
from app.repositories.category_repository import CategoriaRepository

router = APIRouter()

@router.get("/categories")
def listar_categorias():
    categorias = CategoriaRepository.obtener_todas()
    return [c.to_dict() for c in categorias]
