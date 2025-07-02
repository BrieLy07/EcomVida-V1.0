const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');
const authService = require('../services/authService');

const login = async (req, res) => {
  const { correo, contraseña } = req.body;
  try {
    const usuario = await authService.obtenerUsuarioPorCorreo(correo);
    if (!usuario) return res.status(401).json({ error: 'Credenciales inválidas' });

    const match = await bcrypt.compare(contraseña, usuario.contraseña);
    if (!match) return res.status(401).json({ error: 'Credenciales inválidas' });

    const token = jwt.sign(
      { id: usuario.id, correo: usuario.correo, usuario: usuario.usuario },
      process.env.JWT_SECRET,
      { expiresIn: '1h' }
    );

    res.json({ token });
  } catch (err) {
    console.error('Error al iniciar sesión:', err);
    res.status(500).json({ error: 'Error interno al iniciar sesión' });
  }
};

module.exports = login;
