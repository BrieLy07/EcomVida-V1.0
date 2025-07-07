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

@router.put("/{id}")
def actualizar_pedido(id: int, datos: dict, db: Session = Depends(get_db)):
    actualizado = order_repository.actualizar(db, id, datos)
    if not actualizado:
        raise HTTPException(status_code=404, detail="Pedido no encontrado")
    return actualizado
