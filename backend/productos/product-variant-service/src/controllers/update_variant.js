const Variante = require("../models/Variante");

module.exports = async (req, res) => {
  const variante = await Variante.findByIdAndUpdate(req.params.id, req.body, { new: true });
  if (!variante) {
    return res.status(404).json({ mensaje: "Variante no encontrada" });
  }
  res.json(variante);
};
