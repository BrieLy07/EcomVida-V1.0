const sequelize = require('./src/config/db'); // Ruta hacia db.js
require('dotenv').config();

// Importa el modelo para que se registre en Sequelize
require('./src/models/User');

const iniciarBD = async () => {
  try {
    await sequelize.authenticate();
    console.log('✅ Conexión a la base de datos exitosa.');
    await sequelize.sync(); // Crea las tablas si no existen
    console.log('✅ Tablas sincronizadas.');
    process.exit(0);
  } catch (error) {
    console.error('❌ Error al inicializar la base de datos:', error);
    process.exit(1);
  }
};

iniciarBD();
