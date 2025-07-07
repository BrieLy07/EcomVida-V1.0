const Delivery = require("../models/Delivery");

module.exports = async (req, res) => {
  try {
    const entrega = await Delivery.findById(req.params.id);
    if (!entrega) return res.status(404).json({ error: "Entrega no encontrada" });
    res.json(entrega);
  } catch (err) {
    res.status(500).json({ error: "Error al obtener entrega" });
  }
};
