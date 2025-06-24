<?php
require_once __DIR__ . '/../vendor/autoload.php';

use Aws\DynamoDb\DynamoDbClient;

function conectarDynamoDB() {
    return new DynamoDbClient([
        'region' => 'us-west-2',
        'version' => 'latest',
        'endpoint' => getenv('DYNAMODB_ENDPOINT'),
        'credentials' => [
            'key'    => 'fakeMyKeyId',
            'secret' => 'fakeSecretAccessKey'
        ]
    ]);
}
