from fastapi import FastAPI
from app.routes import router
import os
from dotenv import load_dotenv

load_dotenv()

app = FastAPI()
app.include_router(router)

@app.get("/")
def raiz():
    return {"mensaje": "Microservicio de marcas activo"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=int(os.getenv("PORT", 3008)))
