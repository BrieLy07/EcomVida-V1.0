from fastapi import APIRouter, HTTPException, Depends
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

@router.delete("/returns/{id}")
def eliminar_devolucion(id: int, db: Session = Depends(get_db)):
    devolucion = db.query(Return).get(id)
    if not devolucion:
        raise HTTPException(status_code=404, detail="Devolución no encontrada")

    db.delete(devolucion)
    db.commit()
    return {"mensaje": "Devolución eliminada"}
