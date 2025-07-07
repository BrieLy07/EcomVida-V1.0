import strawberry
from typing import Optional
from bson import ObjectId
from datetime import datetime
from app.database.connection import db
from app.graphql.queries.review_queries import ReviewType

@strawberry.type
class MutationReview:

    @strawberry.mutation
    def crearResena(
        self,
        producto_id: str,
        usuario_id: str,
        titulo: str,
        contenido: str,
        rating: int
    ) -> ReviewType:
        nueva_resena = {
            "producto_id": producto_id,
            "usuario_id": usuario_id,
            "titulo": titulo,
            "contenido": contenido,
            "rating": rating,
            "fecha_creacion": datetime.utcnow()
        }
        result = db["reviews"].insert_one(nueva_resena)
        nueva_resena["_id"] = result.inserted_id
        return ReviewType(
            id=str(result.inserted_id),
            **{k: nueva_resena[k] for k in nueva_resena if k != "_id"}
        )

    @strawberry.mutation
    def editarResena(
        self,
        id: strawberry.ID,
        titulo: Optional[str] = None,
        contenido: Optional[str] = None,
        rating: Optional[int] = None
    ) -> Optional[ReviewType]:
        campos = {}
        if titulo: campos["titulo"] = titulo
        if contenido: campos["contenido"] = contenido
        if rating is not None: campos["rating"] = rating

        if not campos:
            return None

        db["reviews"].update_one(
            {"_id": ObjectId(id)},
            {"$set": campos}
        )

        resena = db["reviews"].find_one({"_id": ObjectId(id)})
        if not resena:
            return None

        return ReviewType(
            id=str(resena["_id"]),
            producto_id=resena["producto_id"],
            usuario_id=resena["usuario_id"],
            titulo=resena["titulo"],
            contenido=resena["contenido"],
            rating=resena["rating"],
            fecha_creacion=resena["fecha_creacion"]
        )

    @strawberry.mutation
    def eliminarResena(self, id: strawberry.ID) -> bool:
        result = db["reviews"].delete_one({"_id": ObjectId(id)})
        return result.deleted_count == 1
