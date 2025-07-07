from fastapi import APIRouter, HTTPException
from app.repositories.marca_repository import MarcaRepository

router = APIRouter()

@router.delete("/brands/{marca_id}")
def eliminar_marca(marca_id: int):
    actual = MarcaRepository.obtener_por_id(marca_id)
    if not actual:
        raise HTTPException(status_code=404, detail="Marca no encontrada")

    MarcaRepository.eliminar(marca_id)
    return {"mensaje": "Marca eliminada"}
