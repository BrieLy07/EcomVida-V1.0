from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from . import schemas, crud, database

router = APIRouter()

def get_db():
    db = database.SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.post("/productos/", response_model=schemas.ProductoRespuesta)
def crear_producto(producto: schemas.ProductoCrear, db: Session = Depends(get_db)):
    return crud.crear_producto(db, producto)

@router.get("/productos/", response_model=list[schemas.ProductoRespuesta])
def listar_productos(db: Session = Depends(get_db)):
    return crud.obtener_productos(db)

@router.get("/productos/{producto_id}", response_model=schemas.ProductoRespuesta)
def obtener_producto(producto_id: int, db: Session = Depends(get_db)):
    producto = crud.obtener_producto_por_id(db, producto_id)
    if not producto:
        raise HTTPException(status_code=404, detail="Producto no encontrado")
    return producto
