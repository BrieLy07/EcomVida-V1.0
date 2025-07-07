const bcrypt = require('bcrypt');
const authService = require('../services/authService');

const register = async (req, res) => {
  const { nombre, apellido, usuario, correo, numero, contraseña } = req.body;
  console.log('Datos recibidos:', req.body); // 👈 Añade esto

  try {
    const hashedPassword = await bcrypt.hash(contraseña, 10);
    const nuevoUsuario = await authService.crearUsuario({ nombre, apellido, usuario, correo, numero, contraseña: hashedPassword });
    res.status(201).json(nuevoUsuario);
  } catch (err) {
    console.error('Error al registrar:', err); // 👈 Añade esto también
    res.status(500).json({ error: 'Error al registrar usuario' });
  }
};

module.exports = register;
