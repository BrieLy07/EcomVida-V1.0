from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from app.database.conexion import SessionLocal
from app.repositories import order_repository

router = APIRouter()

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.delete("/{id}")
def eliminar_pedido(id: int, db: Session = Depends(get_db)):
    eliminado = order_repository.eliminar(db, id)
    if not eliminado:
        raise HTTPException(status_code=404, detail="Pedido no encontrado")
    return {"mensaje": "Pedido eliminado correctamente"}
