from fastapi import APIRouter, HTTPException
from app.models.categoria import Categoria
from app.repositories.category_repository import CategoriaRepository

router = APIRouter()

@router.put("/categories/{categoria_id}")
def actualizar_categoria(categoria_id: int, data: dict):
    actual = CategoriaRepository.obtener_por_id(categoria_id)
    if not actual:
        raise HTTPException(status_code=404, detail="Categoría no encontrada")

    nueva = Categoria(None, data["nombre"], data["descripcion"])
    CategoriaRepository.actualizar(categoria_id, nueva)
    return {"mensaje": "Categoría actualizada"}
