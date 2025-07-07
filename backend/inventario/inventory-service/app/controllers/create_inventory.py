from fastapi import Depends
from pydantic import BaseModel
from sqlalchemy.orm import Session
from app.database.mysql_connection import SessionLocal
from app.repository.inventario_repository import InventarioRepository

class InventarioIn(BaseModel):
    producto: str
    cantidad: int
    estado: str = "disponible"

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

def crear(data: InventarioIn, db: Session = Depends(get_db)):
    repo = InventarioRepository(db)
    return repo.crear(data.dict())
