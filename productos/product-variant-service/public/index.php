<?php

require __DIR__ . '/../vendor/autoload.php';

use Slim\Factory\AppFactory;
use App\VariantController;

$app = AppFactory::create();
$app->addBodyParsingMiddleware();

$controller = new VariantController();

$app->post('/variantes', [$controller, 'crear']);
$app->get('/variantes', [$controller, 'listar']);
$app->get('/variantes/{id}', [$controller, 'obtener']);

$app->run();
