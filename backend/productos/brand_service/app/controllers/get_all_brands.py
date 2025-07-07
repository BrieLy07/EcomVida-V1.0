from fastapi import APIRouter
from app.repositories.marca_repository import MarcaRepository

router = APIRouter()

@router.get("/brands")
def listar_marcas():
    marcas = MarcaRepository.obtener_todas()
    return [m.to_dict() for m in marcas]
