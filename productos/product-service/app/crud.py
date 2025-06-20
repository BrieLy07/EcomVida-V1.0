from sqlalchemy.orm import Session
from . import models, schemas

def crear_producto(db: Session, producto: schemas.ProductoCrear):
    nuevo = models.Producto(**producto.dict())
    db.add(nuevo)
    db.commit()
    db.refresh(nuevo)
    return nuevo

def obtener_productos(db: Session):
    return db.query(models.Producto).all()

def obtener_producto_por_id(db: Session, producto_id: int):
    return db.query(models.Producto).filter(models.Producto.id == producto_id).first()
