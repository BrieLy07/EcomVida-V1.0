from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from app.database.conexion import SessionLocal
from app.models.payment import Pago

router = APIRouter()

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.get("/{id}")
def obtener_pago(id: int, db: Session = Depends(get_db)):
    pago = db.query(Pago).filter(Pago.id == id).first()
    if not pago:
        raise HTTPException(status_code=404, detail="Pago no encontrado")
    return pago
