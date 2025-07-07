// start.js
const { initDB } = require('./db'); // ← esto es lo que fallaba
const app = require('./app');

const iniciar = async () => {
  await initDB(); // crea tabla si no existe
  app.listen(3001, () => {
    console.log('Servidor corriendo en puerto 3001');
  });
};

iniciar();
