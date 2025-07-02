const app = require('./app');
const PORT = process.env.PORT || 3001;

const initDB = require('./database/initDB');

initDB.then(() => {
  app.listen(PORT, () => {
    console.log(`Auth Service corriendo en puerto ${PORT}`);
  });
});
