from fastapi import APIRouter, HTTPException
from app.repositories.marca_repository import MarcaRepository

router = APIRouter()

@router.get("/brands/{marca_id}")
def obtener_marca(marca_id: int):
    marca = MarcaRepository.obtener_por_id(marca_id)
    if not marca:
        raise HTTPException(status_code=404, detail="Marca no encontrada")
    return marca.to_dict()
