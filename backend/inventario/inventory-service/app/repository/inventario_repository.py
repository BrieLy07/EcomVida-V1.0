from app.models.inventario import Inventario
from sqlalchemy.orm import Session

class InventarioRepository:

    def __init__(self, db: Session):
        self.db = db

    def listar_todos(self):
        return self.db.query(Inventario).all()

    def obtener_por_id(self, id: int):
        return self.db.query(Inventario).filter(Inventario.id == id).first()

    def crear(self, data):
        nuevo = Inventario(**data)
        self.db.add(nuevo)
        self.db.commit()
        self.db.refresh(nuevo)
        return nuevo

    def actualizar(self, id: int, data):
        inventario = self.obtener_por_id(id)
        if not inventario:
            return None
        for campo, valor in data.items():
            setattr(inventario, campo, valor)
        self.db.commit()
        return inventario

    def eliminar(self, id: int):
        inventario = self.obtener_por_id(id)
        if not inventario:
            return None
        self.db.delete(inventario)
        self.db.commit()
        return inventario
