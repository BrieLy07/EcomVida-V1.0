const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');
const authService = require('../services/authService');

const register = async (req, res) => {
  const { nombre, apellido, usuario, correo, numero, contraseña } = req.body;

  try {
    const hashedPassword = await bcrypt.hash(contraseña, 10);
    const nuevoUsuario = await authService.crearUsuario({
      nombre, apellido, usuario, correo, numero, contraseña: hashedPassword
    });

    // Generar token tras registro
    const token = jwt.sign(
      { id: nuevoUsuario.id, correo: nuevoUsuario.correo, usuario: nuevoUsuario.usuario },
      process.env.JWT_SECRET,
      { expiresIn: '1h' }
    );

    res.status(201).json({ token });
  } catch (err) {
    console.error('Error al registrar:', err);
    res.status(500).json({ error: 'Error al registrar usuario' });
  }
};

module.exports = register;
