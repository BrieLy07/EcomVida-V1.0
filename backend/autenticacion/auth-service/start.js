const initDB = require('./database/initDB');
const app = require('./app');

const PORT = process.env.PORT || 3001;

const iniciar = async () => {
  await initDB(); // crea tabla si no existe
  app.listen(PORT, () => {
    console.log(`Auth Service corriendo en puerto ${PORT}`);
  });
};

iniciar();
