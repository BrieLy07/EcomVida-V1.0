const db = require('../models/db');
const UserFactory = require('../factories/userFactory');

async function register(req, res) {
  try {
    const nuevoUsuario = await UserFactory.createUser(req.body);

    const query = `
      INSERT INTO usuarios (nombre, apellido, usuario, correo, numero, rol, password)
      VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;
    `;
    const values = [
      nuevoUsuario.nombre,
      nuevoUsuario.apellido,
      nuevoUsuario.usuario,
      nuevoUsuario.correo,
      nuevoUsuario.numero,
      nuevoUsuario.rol,
      nuevoUsuario.password
    ];

    const result = await db.query(query, values);
    const usuarioId = result.rows[0].id;

    res.status(201).json({ mensaje: 'Usuario registrado con éxito', id: usuarioId });
  } catch (error) {
    console.error('Error en registro:', error);
    res.status(500).json({ error: 'Error al registrar usuario' });
  }
}

module.exports = { register };
