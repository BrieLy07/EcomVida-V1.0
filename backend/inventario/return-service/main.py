from fastapi import FastAPI
from app.models.return_model import Base
from app.database.connection import engine
from app.controllers import (
    get_all_returns,
    get_return_by_id,
    create_return,
    update_return,
    delete_return,
)

Base.metadata.create_all(bind=engine)

app = FastAPI()

app.include_router(get_all_returns.router)
app.include_router(get_return_by_id.router)
app.include_router(create_return.router)
app.include_router(update_return.router)
app.include_router(delete_return.router)
