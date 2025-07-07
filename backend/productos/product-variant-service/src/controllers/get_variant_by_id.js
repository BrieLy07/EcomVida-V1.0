const Variante = require("../models/Variante");

module.exports = async (req, res) => {
  const variante = await Variante.findById(req.params.id);
  if (!variante) {
    return res.status(404).json({ mensaje: "Variante no encontrada" });
  }
  res.json(variante);
};
