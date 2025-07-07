const mongoose = require("mongoose");

const DeliverySchema = new mongoose.Schema({
  pedidoId: { type: String, required: true },
  direccion: { type: String, required: true },
  estado: {
    type: String,
    enum: ["pendiente", "en_transito", "entregado", "cancelado"],
    default: "pendiente"
  },
  fechaEntrega: { type: Date },
  observaciones: { type: String }
}, { timestamps: true });

module.exports = mongoose.model("Delivery", DeliverySchema);
