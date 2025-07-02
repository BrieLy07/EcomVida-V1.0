const pool = require('./connection');

const crearTabla = async () => {
  const sql = `
    CREATE TABLE IF NOT EXISTS usuarios (
      id SERIAL PRIMARY KEY,
      nombre VARCHAR(100) NOT NULL,
      apellido VARCHAR(100) NOT NULL,
      usuario VARCHAR(100) NOT NULL UNIQUE,
      correo VARCHAR(100) NOT NULL UNIQUE,
      numero VARCHAR(20),
      contraseña TEXT NOT NULL,
      creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
  `;
  try {
    await pool.query(sql);
    console.log("✅ Tabla 'usuarios' verificada/creada.");
  } catch (err) {
    console.error("❌ Error creando la tabla:", err);
  } finally {
    pool.end();
  }
};

crearTabla();
