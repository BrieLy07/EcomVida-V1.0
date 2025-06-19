<?php
require_once "dynamo.php";
use Firebase\JWT\JWT;
use Firebase\JWT\Key;

// Obtener JWT del header
function getUsuarioIdDesdeJWT() {
    if (!isset($_SERVER['HTTP_AUTHORIZATION'])) return null;

    $header = $_SERVER['HTTP_AUTHORIZATION'];
    if (!str_starts_with($header, 'Bearer ')) return null;

    $token = trim(str_replace('Bearer', '', $header));
    try {
        $payload = JWT::decode($token, new Key($_ENV['JWT_SECRET'], 'HS256'));
        return $payload->id;
    } catch (Exception $e) {
        return null;
    }
}

// Procesar solicitud
$method = $_SERVER['REQUEST_METHOD'];
$uri = $_SERVER['REQUEST_URI'];

if ($uri === '/webhook' && $method === 'POST') {
    $usuarioId = getUsuarioIdDesdeJWT();
    if (!$usuarioId) {
        http_response_code(401);
        echo json_encode(['error' => 'No autorizado']);
        exit;
    }

    $input = json_decode(file_get_contents("php://input"), true);
    $ok = guardarPreferencias($usuarioId, $input);
    echo json_encode(['ok' => $ok]);
    exit;
}

if ($uri === '/preferencias' && $method === 'GET') {
    $usuarioId = getUsuarioIdDesdeJWT();
    if (!$usuarioId) {
        http_response_code(401);
        echo json_encode(['error' => 'No autorizado']);
        exit;
    }

    $data = obtenerPreferencias($usuarioId);
    echo json_encode(['preferencias' => $data]);
    exit;
}

// Ruta no encontrada
http_response_code(404);
echo json_encode(['error' => 'Ruta no encontrada']);
