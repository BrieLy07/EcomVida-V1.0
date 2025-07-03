from pymongo import MongoClient
import os

class MongoDB:
    _cliente = None

    @classmethod
    def obtener_cliente(cls):
        if cls._cliente is None:
            uri = os.getenv("MONGO_URI")
            cls._cliente = MongoClient(uri)
        return cls._cliente

    @classmethod
    def obtener_bd(cls):
        cliente = cls.obtener_cliente()
        return cliente[os.getenv("MONGO_DB")]
