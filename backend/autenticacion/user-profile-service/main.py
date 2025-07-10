from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from routes.get_profile import router as get_router
from routes.update_profile import router as update_router

app = FastAPI()

# 👇 Agrega esto para permitir peticiones desde el frontend
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Puedes restringirlo a ["http://localhost:5500"] si deseas
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Rutas
app.include_router(get_router)
app.include_router(update_router)

@app.get("/")
def raiz():
    return {"mensaje": "user-profile-service activo"}
