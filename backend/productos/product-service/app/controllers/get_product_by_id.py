from fastapi import APIRouter, HTTPException
from app.repositories.producto_repository import ProductoRepository

router = APIRouter()

@router.get("/products/{producto_id}")
def obtener_producto(producto_id: int):
    producto = ProductoRepository.obtener_por_id(producto_id)
    if not producto:
        raise HTTPException(status_code=404, detail="Producto no encontrado")
    return producto.to_dict()
