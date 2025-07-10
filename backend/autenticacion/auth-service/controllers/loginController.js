const db = require('../models/db');
const bcrypt = require('bcryptjs');
const { generarToken } = require('../utils/jwt');

async function login(req, res) {
  const { usuario, password } = req.body;

  try {
    const query = 'SELECT * FROM usuarios WHERE usuario = $1';
    const result = await db.query(query, [usuario]);

    if (result.rows.length === 0) {
      return res.status(401).json({ error: 'Usuario no encontrado' });
    }

    const user = result.rows[0];
    const esValido = await bcrypt.compare(password, user.password);

    if (!esValido) {
      return res.status(401).json({ error: 'Contraseña incorrecta' });
    }

    const token = generarToken(user.id, user.rol);
    res.json({ mensaje: 'Login exitoso', token });
  } catch (error) {
    console.error('Error en login:', error);
    res.status(500).json({ error: 'Error al iniciar sesión' });
  }
}

module.exports = { login };
