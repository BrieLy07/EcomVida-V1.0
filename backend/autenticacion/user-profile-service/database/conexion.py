import os
from pymongo import MongoClient
from dotenv import load_dotenv

load_dotenv()

class ConexionMongo:
    _cliente = None

    @classmethod
    def obtener_conexion(cls):
        if cls._cliente is None:
            cls._cliente = MongoClient(os.getenv("MONGO_URI"))
        return cls._cliente[os.getenv("MONGO_DB")]
