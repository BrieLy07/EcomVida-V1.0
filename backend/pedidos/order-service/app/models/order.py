from sqlalchemy import Column, Integer, String, Float, DateTime
from sqlalchemy.ext.declarative import declarative_base
import datetime

Base = declarative_base()

class Pedido(Base):
    __tablename__ = "pedidos"

    id = Column(Integer, primary_key=True, index=True)
    cliente = Column(String, nullable=False)
    producto = Column(String, nullable=False)
    cantidad = Column(Integer, nullable=False)
    total = Column(Float, nullable=False)
    estado = Column(String, default="pendiente")
    creado_en = Column(DateTime, default=datetime.datetime.utcnow)
