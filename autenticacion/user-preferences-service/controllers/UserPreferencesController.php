<?php
require_once __DIR__ . '/../services/UserPreferencesService.php';

class UserPreferencesController {
    public function recibirPreferencias($data) {
        $servicio = new UserPreferencesService();
        $servicio->guardarPreferencias($data);
        echo json_encode(['mensaje' => 'Preferencias guardadas']);
    }
}
