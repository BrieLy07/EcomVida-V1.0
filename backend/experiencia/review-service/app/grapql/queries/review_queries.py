import strawberry
from typing import List, Optional
from app.models.review_model import Review
from app.database.connection import db
from bson import ObjectId
from datetime import datetime

@strawberry.type
class ReviewType:
    id: strawberry.ID
    producto_id: str
    usuario_id: str
    titulo: str
    contenido: str
    rating: int
    fecha_creacion: datetime

@strawberry.type
class QueryReview:

    @strawberry.field
    def obtenerResenas(self) -> List[ReviewType]:
        reviews = db["reviews"].find()
        return [ReviewType(
            id=str(r["_id"]),
            producto_id=r["producto_id"],
            usuario_id=r["usuario_id"],
            titulo=r["titulo"],
            contenido=r["contenido"],
            rating=r["rating"],
            fecha_creacion=r["fecha_creacion"]
        ) for r in reviews]

    @strawberry.field
    def resenaPorId(self, id: strawberry.ID) -> Optional[ReviewType]:
        r = db["reviews"].find_one({"_id": ObjectId(id)})
        if not r:
            return None
        return ReviewType(
            id=str(r["_id"]),
            producto_id=r["producto_id"],
            usuario_id=r["usuario_id"],
            titulo=r["titulo"],
            contenido=r["contenido"],
            rating=r["rating"],
            fecha_creacion=r["fecha_creacion"]
        )

    @strawberry.field
    def resenasPorProducto(self, producto_id: str) -> List[ReviewType]:
        reviews = db["reviews"].find({"producto_id": producto_id})
        return [ReviewType(
            id=str(r["_id"]),
            producto_id=r["producto_id"],
            usuario_id=r["usuario_id"],
            titulo=r["titulo"],
            contenido=r["contenido"],
            rating=r["rating"],
            fecha_creacion=r["fecha_creacion"]
        ) for r in reviews]
