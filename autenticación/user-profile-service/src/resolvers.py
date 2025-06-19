import graphene
from graphene import String, ObjectType, Field
from auth import obtener_usuario_autenticado
from database import coleccion_perfiles

# Tipo de dato para el perfil
class Perfil(graphene.ObjectType):
    nombre = String()
    apellido = String()
    correo = String()
    numero = String()
    foto = String()

# Consulta: obtener perfil
class Query(ObjectType):
    getPerfil = Field(Perfil)

    def resolve_getPerfil(root, info):
        user_id = obtener_usuario_autenticado()
        if not user_id:
            raise Exception("No autorizado")

        perfil = coleccion_perfiles.find_one({"_id": user_id})
        if not perfil:
            raise Exception("Perfil no encontrado")

        return Perfil(
            nombre=perfil.get("nombre"),
            apellido=perfil.get("apellido"),
            correo=perfil.get("correo"),
            numero=perfil.get("numero"),
            foto=perfil.get("foto")
        )

# Mutación: actualizar perfil
class ActualizarPerfil(graphene.Mutation):
    class Arguments:
        nombre = String()
        apellido = String()
        correo = String()
        numero = String()
        foto = String()

    ok = graphene.Boolean()

    def mutate(root, info, nombre=None, apellido=None, correo=None, numero=None, foto=None):
        user_id = obtener_usuario_autenticado()
        if not user_id:
            raise Exception("No autorizado")

        coleccion_perfiles.update_one(
            {"_id": user_id},
            {"$set": {
                "nombre": nombre,
                "apellido": apellido,
                "correo": correo,
                "numero": numero,
                "foto": foto
            }},
            upsert=True
        )
        return ActualizarPerfil(ok=True)

class Mutation(ObjectType):
    actualizarPerfil = ActualizarPerfil.Field()
