from fastapi import APIRouter, Depends
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

@router.get("")
def listar_pedidos(db: Session = Depends(get_db)):
    return order_repository.obtener_todos(db)
