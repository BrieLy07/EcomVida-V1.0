const Variante = require("../models/Variante");

module.exports = async (req, res) => {
  const variantes = await Variante.find();
  res.json(variantes);
};
