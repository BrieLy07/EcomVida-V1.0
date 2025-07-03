from app.db.conexion import MongoDB
from bson.objectid import ObjectId

def actualizar_usuario(id_usuario: str, datos: dict):
    db = MongoDB.obtener_bd()
    resultado = db.usuarios.update_one(
        {"_id": ObjectId(id_usuario)},
        {"$set": datos}
    )
    return resultado.modified_count > 0
