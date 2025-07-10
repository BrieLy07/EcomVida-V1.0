CREATE TABLE IF NOT EXISTS usuarios (
  id SERIAL PRIMARY KEY,
  nombre VARCHAR(100),
  apellido VARCHAR(100),
  usuario VARCHAR(50) UNIQUE,
  correo VARCHAR(100) UNIQUE,
  numero VARCHAR(15),
  rol VARCHAR(20),
  password TEXT
);
