from fastapi import APIRouter
from app.repositories.producto_repository import ProductoRepository

router = APIRouter()

@router.get("/products")
def obtener_productos():
    productos = ProductoRepository.obtener_todos()
    return [p.to_dict() for p in productos]
