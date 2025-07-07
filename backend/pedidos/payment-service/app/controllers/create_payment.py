from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from app.database.conexion import SessionLocal
from app.models.payment import Pago
from app.templates.procesadores.tarjeta import PagoTarjeta

router = APIRouter()

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.post("")
def crear_pago(pago: dict, db: Session = Depends(get_db)):
    procesador = PagoTarjeta()
    resultado = procesador.procesar(pago)

    if resultado["exito"]:
        nuevo_pago = Pago(
            metodo="tarjeta",
            monto=pago["monto"],
            estado="aprobado",
            orden_id=pago["orden_id"]
        )
        db.add(nuevo_pago)
        db.commit()
        db.refresh(nuevo_pago)
        return nuevo_pago
    else:
        raise HTTPException(status_code=400, detail="Pago rechazado")
