from fastapi import APIRouter, HTTPException
from app.models.marca import Marca
from app.repositories.marca_repository import MarcaRepository

router = APIRouter()

@router.put("/brands/{marca_id}")
def actualizar_marca(marca_id: int, data: dict):
    actual = MarcaRepository.obtener_por_id(marca_id)
    if not actual:
        raise HTTPException(status_code=404, detail="Marca no encontrada")

    nueva = Marca(None, data["nombre"], data["descripcion"])
    MarcaRepository.actualizar(marca_id, nueva)
    return {"mensaje": "Marca actualizada"}
