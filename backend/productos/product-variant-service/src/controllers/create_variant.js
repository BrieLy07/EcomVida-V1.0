const Variante = require("../models/Variante");
const VarianteBuilder = require("../builders/VarianteBuilder");

module.exports = async (req, res) => {
  const { color, talla, precioExtra, disponible } = req.body;
  const { productId } = req.params;

  const builder = new VarianteBuilder()
    .setProductoId(productId)
    .setColor(color)
    .setTalla(talla)
    .setPrecioExtra(precioExtra)
    .setDisponible(disponible);

  const nuevaVariante = new Variante(builder.build());
  await nuevaVariante.save();
  res.status(201).json(nuevaVariante);
};
