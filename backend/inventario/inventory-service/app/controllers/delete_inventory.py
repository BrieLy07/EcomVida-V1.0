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

def eliminar(id: int, db: Session = Depends(get_db)):
    repo = InventarioRepository(db)
    eliminado = repo.eliminar(id)
    if eliminado:
        return {"mensaje": "Eliminado"}
    return {"error": "No encontrado"}
