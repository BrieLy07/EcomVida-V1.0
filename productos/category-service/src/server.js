const app = require("./app");
const sequelize = require("./config/database");
const Categoria = require("./models/category");

const PORT = process.env.PORT || 3002;

const iniciarServidor = async () => {
  let conectado = false;

  while (!conectado) {
    try {
      await sequelize.authenticate();
      console.log("✅ Conexión a MySQL establecida.");
      conectado = true;
    } catch (error) {
      console.log("⏳ Esperando base de datos...");
      await new Promise(res => setTimeout(res, 3000)); // espera 3 segundos
    }
  }

  await sequelize.sync();
  console.log("📦 Base de datos sincronizada");

  app.listen(PORT, () => {
    console.log(`🚀 Servidor en http://localhost:${PORT}`);
  });
};

iniciarServidor();
