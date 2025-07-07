from sqlalchemy import Column, Integer, String, Float
from app.database.conexion import Base

class Pago(Base):
    __tablename__ = "pagos"

    id = Column(Integer, primary_key=True, index=True)
    metodo = Column(String, nullable=False)
    monto = Column(Float, nullable=False)
    estado = Column(String, default="pendiente")
    orden_id = Column(Integer, nullable=False)
