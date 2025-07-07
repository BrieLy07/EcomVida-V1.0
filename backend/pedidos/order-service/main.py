from fastapi import FastAPI
from dotenv import load_dotenv
import os

load_dotenv()

app = FastAPI(title="Order Service")

# Rutas
from app.controllers.get_orders import router as get_orders_router
from app.controllers.get_order_by_id import router as get_order_router
from app.controllers.create_order import router as create_order_router
from app.controllers.update_order import router as update_order_router
from app.controllers.delete_order import router as delete_order_router

app.include_router(get_orders_router, prefix="/orders")
app.include_router(get_order_router, prefix="/orders")
app.include_router(create_order_router, prefix="/orders")
app.include_router(update_order_router, prefix="/orders")
app.include_router(delete_order_router, prefix="/orders")
