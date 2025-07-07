from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session
from app.database.conexion import SessionLocal
from app.repositories import order_repository
from app.models.order import Pedido

router = APIRouter()

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.post("")
def crear_pedido(pedido: Pedido, db: Session = Depends(get_db)):
    return order_repository.crear(db, pedido)
