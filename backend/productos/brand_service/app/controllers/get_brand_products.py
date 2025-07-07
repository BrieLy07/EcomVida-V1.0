from fastapi import APIRouter, HTTPException
from app.repositories.marca_repository import MarcaRepository

router = APIRouter()

@router.get("/brands/{marca_id}/products")
def obtener_productos_marca(marca_id: int):
    marca = MarcaRepository.obtener_por_id(marca_id)
    if not marca:
        raise HTTPException(status_code=404, detail="Marca no encontrada")

    productos = MarcaRepository.obtener_productos_por_marca(marca_id)
    return {
        "marca": marca.to_dict(),
        "productos": [p.to_dict() for p in productos]
    }
