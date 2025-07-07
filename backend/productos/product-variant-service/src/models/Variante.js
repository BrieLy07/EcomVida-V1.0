const mongoose = require("mongoose");

const varianteSchema = new mongoose.Schema({
  productoId: { type: String, required: true },
  color: String,
  talla: String,
  precioExtra: Number,
  disponible: Boolean
});

module.exports = mongoose.model("Variante", varianteSchema);
