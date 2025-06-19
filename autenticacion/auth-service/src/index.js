const express = require('express');
const app = express();
const sequelize = require('./config/db');
const authRoutes = require('./routes/authRoutes');
require('dotenv').config();

app.use(express.json());
app.use('/api/auth', authRoutes);

// Reintentar conexión a base de datos
const iniciarServidor = async (intentos = 10) => {
  while (intentos > 0) {
    try {
      await sequelize.authenticate();
      await sequelize.sync();
      console.log('Base de datos conectada');
      app.listen(process.env.PORT, () => {
        console.log(`Auth Service corriendo en puerto ${process.env.PORT}`);
      });
      break;
    } catch (error) {
      console.log('Esperando base de datos... reintentando en 3 segundos');
      intentos--;
      await new Promise((res) => setTimeout(res, 3000));
    }
  }

  if (intentos === 0) {
    console.error('No se pudo conectar a la base de datos. Verifica configuración.');
    process.exit(1);
  }
};

iniciarServidor();
