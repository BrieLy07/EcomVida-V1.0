from fastapi import FastAPI
from app.models.inventario import Base
from app.database.mysql_connection import engine
from app.controllers import (
    get_all_inventory,
    get_inventory_by_id,
    create_inventory,
    update_inventory,
    delete_inventory
)

# Crear app
app = FastAPI()

# Crear tablas automáticamente al iniciar
Base.metadata.create_all(bind=engine)

# Rutas del microservicio
app.get("/inventory")(get_all_inventory.listar_todos)
app.get("/inventory/{id}")(get_inventory_by_id.obtener_por_id)
app.post("/inventory")(create_inventory.crear)
app.put("/inventory/{id}")(update_inventory.actualizar)
app.delete("/inventory/{id}")(delete_inventory.eliminar)

# Ruta base de prueba
@app.get("/")
def home():
    return {"mensaje": "Microservicio de Inventario activo"}
