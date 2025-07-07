from fastapi import APIRouter, HTTPException
from app.repositories.category_repository import CategoriaRepository

router = APIRouter()

@router.get("/categories/{categoria_id}/products")
def obtener_productos_categoria(categoria_id: int):
    categoria = CategoriaRepository.obtener_por_id(categoria_id)
    if not categoria:
        raise HTTPException(status_code=404, detail="Categoría no encontrada")

    productos = CategoriaRepository.obtener_productos_por_categoria(categoria_id)
    for p in productos:
        categoria.agregar_producto(p)

    return categoria.to_dict(incluir_productos=True)
