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

def listar_todos(db: Session = Depends(get_db)):
    repo = InventarioRepository(db)
    return repo.listar_todos()
