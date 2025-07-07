from fastapi import APIRouter, HTTPException
from app.repositories.producto_repository import ProductoRepository
from app.models.producto import Producto

router = APIRouter()

@router.put("/products/{producto_id}")
def actualizar_producto(producto_id: int, data: dict):
    producto_existente = ProductoRepository.obtener_por_id(producto_id)
    if not producto_existente:
        raise HTTPException(status_code=404, detail="Producto no encontrado")

    actualizado = Producto(None, data["nombre"], data["descripcion"], data["precio"], data["stock"])
    ProductoRepository.actualizar(producto_id, actualizado)
    return {"mensaje": "Producto actualizado"}
