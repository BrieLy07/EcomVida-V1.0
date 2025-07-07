from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session
from app.database.connection import SessionLocal
from app.models.return_model import Return

router = APIRouter()

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.get("/returns")
def listar_devoluciones(db: Session = Depends(get_db)):
    return db.query(Return).all()
