<?php
require_once __DIR__ . '/../database/connection.php';

class UserPreferencesService {
    public function guardarPreferencias($data) {
        $conexion = conectarDynamoDB();
        $conexion->putItem([
            'TableName' => 'user_preferences',
            'Item' => [
                'usuario' => ['S' => $data['usuario']],
                'preferencias' => ['S' => json_encode($data['preferencias'])]
            ]
        ]);
    }
}
