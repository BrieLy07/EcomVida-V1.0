const Factura = require("../models/Factura");

module.exports = async (req, res) => {
  try {
    const actualizada = await Factura.findByIdAndUpdate(
      req.params.id,
      req.body,
      { new: true }
    );
    if (!actualizada) return res.status(404).json({ error: "No encontrada" });
    res.json(actualizada);
  } catch (err) {
    res.status(500).json({ error: "Error al actualizar factura" });
  }
};
