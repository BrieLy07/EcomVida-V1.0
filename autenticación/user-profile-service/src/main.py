from flask import Flask
from flask_cors import CORS
from flask_graphql import GraphQLView
from schema import schema
import os
from dotenv import load_dotenv

load_dotenv()
PORT = int(os.getenv("PORT", 5000))

app = Flask(__name__)
CORS(app)

app.add_url_rule(
    '/graphql',
    view_func=GraphQLView.as_view(
        'graphql',
        schema=schema,
        graphiql=True  # Habilita interfaz web para pruebas
    )
)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=PORT)
