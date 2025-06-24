<?php
require_once __DIR__ . '/../controllers/UserPreferencesController.php';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $input = json_decode(file_get_contents('php://input'), true);
    (new UserPreferencesController())->recibirPreferencias($input);
}
