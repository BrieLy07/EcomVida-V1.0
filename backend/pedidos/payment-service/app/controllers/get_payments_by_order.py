from fastapi import APIRouter, Depends
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

@router.get("/{order_id}")
def pagos_por_orden(order_id: int, db: Session = Depends(get_db)):
    return db.query(Pago).filter(Pago.orden_id == order_id).all()
