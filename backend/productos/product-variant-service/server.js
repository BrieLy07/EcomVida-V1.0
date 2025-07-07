const express = require("express");
const app = express();
const dotenv = require("dotenv");
const db = require("./src/database");
const routes = require("./src/routes");

dotenv.config();

app.use(express.json());
app.use("/", routes);

const PORT = process.env.PORT || 3009;
app.listen(PORT, () => {
  console.log(`Servidor de variantes corriendo en el puerto ${PORT}`);
});
