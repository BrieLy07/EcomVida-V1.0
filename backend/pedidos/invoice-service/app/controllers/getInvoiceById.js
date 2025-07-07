const Factura = require("../models/Factura");

module.exports = async (req, res) => {
  try {
    const factura = await Factura.findById(req.params.id);
    if (!factura) return res.status(404).json({ error: "No encontrada" });
    res.json(factura);
  } catch (err) {
    res.status(500).json({ error: "Error al obtener factura" });
  }
};
