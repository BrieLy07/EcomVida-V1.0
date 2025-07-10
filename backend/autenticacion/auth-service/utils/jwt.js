const jwt = require('jsonwebtoken');
require('dotenv').config();

function generarToken(usuarioId, rol) {
  return jwt.sign({ id: usuarioId, rol }, process.env.JWT_SECRET, { expiresIn: '2h' });
}

module.exports = { generarToken };
