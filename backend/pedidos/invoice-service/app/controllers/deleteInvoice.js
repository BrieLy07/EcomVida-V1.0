const Factura = require("../models/Factura");

module.exports = async (req, res) => {
  try {
    const eliminada = await Factura.findByIdAndDelete(req.params.id);
    if (!eliminada) return res.status(404).json({ error: "No encontrada" });
    res.json({ mensaje: "Factura eliminada" });
  } catch (err) {
    res.status(500).json({ error: "Error al eliminar factura" });
  }
};
