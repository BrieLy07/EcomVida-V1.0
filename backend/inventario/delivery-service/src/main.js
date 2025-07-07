const express = require("express");
const dotenv = require("dotenv");
const cors = require("cors");
const connectMongo = require("./database/mongo");

// Cargar controladores
const getAll = require("./controllers/get_all_deliveries");
const getById = require("./controllers/get_delivery_by_id");
const create = require("./controllers/create_delivery");
const update = require("./controllers/update_delivery");
const remove = require("./controllers/delete_delivery");

// Configuración
dotenv.config();
connectMongo();

const app = express();
app.use(cors());
app.use(express.json());

// Rutas
app.get("/deliveries", getAll);
app.get("/deliveries/:id", getById);
app.post("/deliveries", create);
app.put("/deliveries/:id", update);
app.delete("/deliveries/:id", remove);

// Iniciar servidor
const PORT = process.env.PORT || 3019;
app.listen(PORT, () => {
  console.log(`📦 delivery-service corriendo en puerto ${PORT}`);
});
