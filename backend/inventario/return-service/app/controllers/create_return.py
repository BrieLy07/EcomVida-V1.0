from fastapi import APIRouter, HTTPException, Depends
from sqlalchemy.orm import Session
from app.database.connection import SessionLocal
from app.models.return_model import Return
from app.handlers.validate_motivo_handler import ValidateMotivoHandler
from app.handlers.validate_estado_handler import ValidateEstadoHandler

router = APIRouter()

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.post("/returns")
def registrar_devolucion(data: dict, db: Session = Depends(get_db)):
    # Cadena de validaciones
    handler = ValidateMotivoHandler()
    handler.set_next(ValidateEstadoHandler())
    handler.handle(data)

    devolucion = Return(**data)
    db.add(devolucion)
    db.commit()
    db.refresh(devolucion)
    return devolucion
