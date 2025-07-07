from fastapi import APIRouter, HTTPException
from app.repositories.category_repository import CategoriaRepository

router = APIRouter()

@router.get("/categories/{categoria_id}")
def obtener_categoria(categoria_id: int):
    categoria = CategoriaRepository.obtener_por_id(categoria_id)
    if not categoria:
        raise HTTPException(status_code=404, detail="Categoría no encontrada")
    return categoria.to_dict()
