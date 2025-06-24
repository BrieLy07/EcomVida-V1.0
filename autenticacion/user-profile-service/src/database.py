import os
from pymongo import MongoClient
from dotenv import load_dotenv

# Cargar variables de entorno
load_dotenv()

# Obtener valores del .env
mongo_uri = os.getenv("MONGO_URI")
db_name = os.getenv("DB_NAME", "perfilusuarios")  # Por defecto usamos "perfilusuarios"

# Conectar a MongoDB Atlas
client = MongoClient(mongo_uri)

# Seleccionar base de datos y colección
db = client[db_name]
coleccion_perfiles = db["perfiles"]
