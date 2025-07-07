from fastapi import APIRouter, HTTPException
from app.repositories.category_repository import CategoriaRepository

router = APIRouter()

@router.delete("/categories/{categoria_id}")
def eliminar_categoria(categoria_id: int):
    actual = CategoriaRepository.obtener_por_id(categoria_id)
    if not actual:
        raise HTTPException(status_code=404, detail="Categoría no encontrada")

    CategoriaRepository.eliminar(categoria_id)
    return {"mensaje": "Categoría eliminada"}
