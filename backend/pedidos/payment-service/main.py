from fastapi import FastAPI
from app.database.conexion import Base, engine
from app.controllers import create_payment, get_payment_by_id, get_payments_by_order

app = FastAPI()

# Crear tablas si no existen
Base.metadata.create_all(bind=engine)

# Incluir rutas
app.include_router(create_payment.router, prefix="/payments")
app.include_router(get_payment_by_id.router, prefix="/payments")
app.include_router(get_payments_by_order.router, prefix="/payments/order")
