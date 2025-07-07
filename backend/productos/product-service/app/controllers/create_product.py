from fastapi import APIRouter
from app.repositories.producto_repository import ProductoRepository
from app.models.producto import Producto

router = APIRouter()

@router.post("/products")
def crear_producto(data: dict):
    producto = Producto(None, data["nombre"], data["descripcion"], data["precio"], data["stock"])
    producto_creado = ProductoRepository.crear(producto)
    return producto_creado.to_dict()
