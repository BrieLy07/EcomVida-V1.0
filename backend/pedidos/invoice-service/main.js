const express = require("express");
const conectarDB = require("./app/database/mongoClient");
require("dotenv").config();

const crearFactura = require("./app/controllers/createInvoice");
const obtenerFactura = require("./app/controllers/getInvoiceById");
const obtenerFacturaPorOrden = require("./app/controllers/getInvoiceByOrder");
const actualizarFactura = require("./app/controllers/updateInvoice");
const eliminarFactura = require("./app/controllers/deleteInvoice");

const app = express();
app.use(express.json());

// Conexión DB
conectarDB();

// Rutas
app.post("/invoices", crearFactura);
app.get("/invoices/:id", obtenerFactura);
app.get("/invoices/order/:orderId", obtenerFacturaPorOrden);
app.put("/invoices/:id", actualizarFactura);
app.delete("/invoices/:id", eliminarFactura);

const PORT = process.env.PORT || 3015;
app.listen(PORT, () => console.log(`Invoice service corriendo en puerto ${PORT}`));
