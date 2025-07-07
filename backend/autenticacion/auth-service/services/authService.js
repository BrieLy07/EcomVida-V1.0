const pool = require('../database/connection');

const crearUsuario = async ({ nombre, apellido, usuario, correo, numero, contraseña }) => {
  const res = await pool.query(
    'INSERT INTO usuarios (nombre, apellido, usuario, correo, numero, contraseña) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *',
    [nombre, apellido, usuario, correo, numero, contraseña]
  );
  return res.rows[0];
};

const obtenerUsuarioPorCorreo = async (correo) => {
  const res = await pool.query('SELECT * FROM usuarios WHERE correo = $1', [correo]);
  return res.rows[0];
};

module.exports = { crearUsuario, obtenerUsuarioPorCorreo };
