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

@router.get("/{id}")
def obtener_pedido(id: int, db: Session = Depends(get_db)):
    pedido = order_repository.obtener_por_id(db, id)
    if not pedido:
        raise HTTPException(status_code=404, detail="Pedido no encontrado")
    return pedido
