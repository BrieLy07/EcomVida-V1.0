const Delivery = require("../models/Delivery");

module.exports = async (req, res) => {
  try {
    await Delivery.findByIdAndDelete(req.params.id);
    res.json({ mensaje: "Entrega eliminada" });
  } catch (err) {
    res.status(500).json({ error: "Error al eliminar entrega" });
  }
};
