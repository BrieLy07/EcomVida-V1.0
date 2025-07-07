from app.db.conexion import MongoDB
from bson.objectid import ObjectId

def obtener_usuario(id: str):
    db = MongoDB.obtener_bd()
    usuario = db.usuarios.find_one({"_id": ObjectId(id)})
    if usuario:
        usuario["_id"] = str(usuario["_id"])
        return usuario
    return None
