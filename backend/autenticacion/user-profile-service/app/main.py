from fastapi import FastAPI
from app.routes.obtener_usuario import router as obtener_router
from app.routes.actualizar_usuario import router as actualizar_router
from app.routes import create_user 
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

app.include_router(obtener_router)
app.include_router(actualizar_router)
app.include_router(create_user.router)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # En pruebas locales
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)