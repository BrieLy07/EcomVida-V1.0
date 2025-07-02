const { Pool } = require('pg');
require('dotenv').config();

// Crear el pool de conexión
const pool = new Pool({
  host: process.env.DB_HOST,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  port: process.env.DB_PORT,
  ssl: {
    rejectUnauthorized: false, // necesario para conexión con RDS
  },
});

// Crear la tabla si no existe
async function initDB() {
  const query = `
    CREATE TABLE IF NOT EXISTS usuarios (
      id SERIAL PRIMARY KEY,
      nombre VARCHAR(100) NOT NULL,
      apellido VARCHAR(100) NOT NULL,
      usuario VARCHAR(50) UNIQUE NOT NULL,
      correo VARCHAR(100) UNIQUE NOT NULL,
      numero VARCHAR(20),
      contraseña TEXT NOT NULL,
      fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
  `;
  try {
    await pool.query(query);
    console.log('✅ Tabla usuarios verificada o creada correctamente.');
  } catch (err) {
    console.error('❌ Error al crear/verificar la tabla:', err.message);
    throw err;
  }
}

module.exports = { pool, initDB };
