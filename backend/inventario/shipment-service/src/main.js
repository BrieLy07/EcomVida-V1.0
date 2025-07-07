const express = require("express");
const dotenv = require("dotenv");
const cors = require("cors");

// Controladores
const getAllShipments = require("./controllers/get_all_shipments");
const getShipmentById = require("./controllers/get_shipment_by_id");
const createShipment = require("./controllers/create_shipment");
const updateShipment = require("./controllers/update_shipment");
const deleteShipment = require("./controllers/delete_shipment");

// Configuración
dotenv.config();
const app = express();
app.use(cors());
app.use(express.json());

// Rutas
app.get("/shipments", getAllShipments);
app.get("/shipments/:id", getShipmentById);
app.post("/shipments", createShipment);
app.put("/shipments/:id", updateShipment);
app.delete("/shipments/:id", deleteShipment);

// Inicio del servidor
const PORT = process.env.PORT || 3018;
app.listen(PORT, () => {
  console.log(`🚚 shipment-service corriendo en el puerto ${PORT}`);
});
