from fastapi import FastAPI
from strawberry.fastapi import GraphQLRouter
from app.graphql.schema import schema
from app.database.connection import connect_db

app = FastAPI()

# Conexión a MongoDB
connect_db()

# Ruteo de GraphQL
graphql_app = GraphQLRouter(schema)
app.include_router(graphql_app, prefix="/graphql")
