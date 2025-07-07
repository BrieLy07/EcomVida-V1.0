const mongoose = require("mongoose");

const FacturaSchema = new mongoose.Schema({
  orderId: { type: String, required: true },
  cliente: { type: String, required: true },
  monto: { type: Number, required: true },
  estado: { type: String, default: "emitida" },
  tipo: { type: String, default: "simple" }, // simple, detallada, etc.
  fecha: { type: Date, default: Date.now }
});

module.exports = mongoose.model("Factura", FacturaSchema);
