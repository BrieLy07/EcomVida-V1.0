from sqlalchemy import Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()

class Inventario(Base):
    __tablename__ = "inventario"

    id = Column(Integer, primary_key=True, index=True)
    producto = Column(String(100), nullable=False)
    cantidad = Column(Integer, nullable=False)
    estado = Column(String(50), default="disponible")
