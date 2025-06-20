<?php

namespace App;

use MongoDB\Client;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

class VariantController {
    private $collection;

    public function __construct() {
        $dotenv = \Dotenv\Dotenv::createImmutable(__DIR__ . '/../');
        $dotenv->load();

        $mongo = new Client($_ENV['MONGO_URI']);
        $this->collection = $mongo->{$_ENV['MONGO_DB']}->variants;
    }

    public function crear(Request $request, Response $response): Response {
        $data = $request->getParsedBody();
        $res = $this->collection->insertOne($data);
        $response->getBody()->write(json_encode(['ok' => true, 'id' => (string)$res->getInsertedId()]));
        return $response->withHeader('Content-Type', 'application/json');
    }

    public function listar(Request $request, Response $response): Response {
        $docs = $this->collection->find()->toArray();
        $response->getBody()->write(json_encode($docs));
        return $response->withHeader('Content-Type', 'application/json');
    }

    public function obtener(Request $request, Response $response, $args): Response {
        $doc = $this->collection->findOne(['_id' => new \MongoDB\BSON\ObjectId($args['id'])]);
        if (!$doc) {
            return $response->withStatus(404)->withHeader('Content-Type', 'application/json')->write(json_encode(['error' => 'No encontrado']));
        }
        $response->getBody()->write(json_encode($doc));
        return $response->withHeader('Content-Type', 'application/json');
    }
}
