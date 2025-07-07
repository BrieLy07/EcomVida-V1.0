const Delivery = require("../models/Delivery");
const getStateInstance = require("../states/StateFactory");

module.exports = async (req, res) => {
  try {
    const entrega = await Delivery.findById(req.params.id);
    if (!entrega) return res.status(404).json({ error: "Entrega no encontrada" });

    const state = getStateInstance(entrega);
    if (!state) return res.status(400).json({ error: "Estado no modificable" });

    // Cambio de estado
    if (req.body.estado === "avanzar") state.avanzar();
    if (req.body.estado === "cancelar") state.cancelar();

    await entrega.save();
    res.json(entrega);
  } catch (err) {
    res.status(500).json({ error: "Error al actualizar entrega" });
  }
};
