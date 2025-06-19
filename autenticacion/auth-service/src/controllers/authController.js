const Usuario = require('../models/User');
const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');

exports.registrar = async (req, res) => {
  try {
    const { nombre, apellido, usuario, correo, numero, contraseña } = req.body;
    const hash = await bcrypt.hash(contraseña, 10);

    const nuevoUsuario = await Usuario.create({
      nombre,
      apellido,
      usuario,
      correo,
      numero,
      contraseña: hash,
    });

    res.status(201).json({ mensaje: 'Usuario registrado', usuario: nuevoUsuario });
  } catch (error) {
    res.status(400).json({ error: error.message });
  }
};

exports.login = async (req, res) => {
  try {
    const { usuario, contraseña } = req.body;
    const user = await Usuario.findOne({ where: { usuario } });

    if (!user) return res.status(404).json({ error: 'Usuario no encontrado' });

    const valido = await bcrypt.compare(contraseña, user.contraseña);
    if (!valido) return res.status(401).json({ error: 'Contraseña incorrecta' });

    const token = jwt.sign({ id: user.id }, process.env.JWT_SECRET, { expiresIn: '1h' });

    res.json({ mensaje: 'Login exitoso', token });
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};
