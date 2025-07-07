import os
from pymongo import MongoClient
from dotenv import load_dotenv

load_dotenv()

client = None
db = None

def connect_db():
    global client, db
    client = MongoClient(os.getenv("MONGO_URI"))
    db = client[os.getenv("DB_NAME")]
