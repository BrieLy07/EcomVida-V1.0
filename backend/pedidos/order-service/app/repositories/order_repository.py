from sqlalchemy.orm import Session
from app.models.order import Pedido

def obtener_todos(db: Session):
    return db.query(Pedido).all()

def obtener_por_id(db: Session, pedido_id: int):
    return db.query(Pedido).filter(Pedido.id == pedido_id).first()

def crear(db: Session, pedido: Pedido):
    db.add(pedido)
    db.commit()
    db.refresh(pedido)
    return pedido

def actualizar(db: Session, pedido_id: int, datos: dict):
    pedido = obtener_por_id(db, pedido_id)
    if not pedido:
        return None
    for clave, valor in datos.items():
        setattr(pedido, clave, valor)
    db.commit()
    db.refresh(pedido)
    return pedido

def eliminar(db: Session, pedido_id: int):
    pedido = obtener_por_id(db, pedido_id)
    if not pedido:
        return None
    db.delete(pedido)
    db.commit()
    return pedido
