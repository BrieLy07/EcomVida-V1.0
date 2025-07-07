from fastapi import Depends
from pydantic import BaseModel
from sqlalchemy.orm import Session
from app.database.mysql_connection import SessionLocal
from app.repository.inventario_repository import InventarioRepository

class InventarioUpdate(BaseModel):
    producto: str = None
    cantidad: int = None
    estado: str = None

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

def actualizar(id: int, data: InventarioUpdate, db: Session = Depends(get_db)):
    repo = InventarioRepository(db)
    return repo.actualizar(id, data.dict(exclude_unset=True))
