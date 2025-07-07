from fastapi import Depends
from sqlalchemy.orm import Session
from app.database.mysql_connection import SessionLocal
from app.repository.inventario_repository import InventarioRepository

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

def obtener_por_id(id: int, db: Session = Depends(get_db)):
    repo = InventarioRepository(db)
    inventario = repo.obtener_por_id(id)
    if inventario:
        return inventario
    return {"error": "No encontrado"}
