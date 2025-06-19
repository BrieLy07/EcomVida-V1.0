<?php
require __DIR__ . '/../vendor/autoload.php';

use Aws\DynamoDb\DynamoDbClient;
use Aws\DynamoDb\Exception\DynamoDbException;

$dotenv = Dotenv\Dotenv::createImmutable(__DIR__ . "/..");
$dotenv->load();

$client = new DynamoDbClient([
    'region'  => $_ENV['AWS_REGION'],
    'version' => 'latest',
    'endpoint' => $_ENV['DYNAMODB_ENDPOINT'],
    'credentials' => [
        'key'    => $_ENV['AWS_ACCESS_KEY_ID'],
        'secret' => $_ENV['AWS_SECRET_ACCESS_KEY'],
    ]
]);

$table = $_ENV['DYNAMODB_TABLE'];

// Crear la tabla si no existe
try {
    $result = $client->listTables();
    if (!in_array($table, $result['TableNames'])) {
        $client->createTable([
            'TableName' => $table,
            'AttributeDefinitions' => [
                ['AttributeName' => 'usuarioId', 'AttributeType' => 'S'],
            ],
            'KeySchema' => [
                ['AttributeName' => 'usuarioId', 'KeyType' => 'HASH'],
            ],
            'ProvisionedThroughput' => [
                'ReadCapacityUnits' => 5,
                'WriteCapacityUnits' => 5
            ]
        ]);
        echo "Tabla '$table' creada en DynamoDB local\n";
    }
} catch (Exception $e) {
    echo "Error al crear la tabla: " . $e->getMessage() . "\n";
}

function guardarPreferencias($usuarioId, $data) {
    global $client, $table;

    $item = [
        'usuarioId' => ['S' => $usuarioId],
        'preferencias' => ['S' => json_encode($data)],
    ];

    try {
        $client->putItem([
            'TableName' => $table,
            'Item' => $item
        ]);
        return true;
    } catch (DynamoDbException $e) {
        return false;
    }
}

function obtenerPreferencias($usuarioId) {
    global $client, $table;

    try {
        $result = $client->getItem([
            'TableName' => $table,
            'Key' => [
                'usuarioId' => ['S' => $usuarioId]
            ]
        ]);

        if (!isset($result['Item'])) return null;

        return json_decode($result['Item']['preferencias']['S'], true);
    } catch (DynamoDbException $e) {
        return null;
    }
}
