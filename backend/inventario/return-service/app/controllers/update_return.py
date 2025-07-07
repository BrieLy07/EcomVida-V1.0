from fastapi import APIRouter, HTTPException, Depends
from sqlalchemy.orm import Session
from app.database.connection import SessionLocal
from app.models.return_model import Return
from app.handlers.validate_estado_handler import ValidateEstadoHandler

router = APIRouter()

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.put("/returns/{id}")
def actualizar_estado(id: int, data: dict, db: Session = Depends(get_db)):
    devolucion = db.query(Return).get(id)
    if not devolucion:
        raise HTTPException(status_code=404, detail="Devolución no encontrada")

    # Validación del estado
    handler = ValidateEstadoHandler()
    handler.handle(data)

    devolucion.estado = data.get("estado", devolucion.estado)
    db.commit()
    db.refresh(devolucion)
    return devolucion
