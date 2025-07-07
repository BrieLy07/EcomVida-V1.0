from fastapi import APIRouter, HTTPException
from app.repositories.producto_repository import ProductoRepository

router = APIRouter()

@router.delete("/products/{producto_id}")
def eliminar_producto(producto_id: int):
    producto = ProductoRepository.obtener_por_id(producto_id)
    if not producto:
        raise HTTPException(status_code=404, detail="Producto no encontrado")

    ProductoRepository.eliminar(producto_id)
    return {"mensaje": "Producto eliminado"}
