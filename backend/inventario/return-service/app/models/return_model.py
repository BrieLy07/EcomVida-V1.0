from sqlalchemy import Column, Integer, String, DateTime
from sqlalchemy.ext.declarative import declarative_base
from datetime import datetime

Base = declarative_base()

class Return(Base):
    __tablename__ = "returns"

    id = Column(Integer, primary_key=True, index=True)
    pedido_id = Column(String, nullable=False)
    motivo = Column(String, nullable=False)
    estado = Column(String, default="pendiente")
    fecha_registro = Column(DateTime, default=datetime.utcnow)
