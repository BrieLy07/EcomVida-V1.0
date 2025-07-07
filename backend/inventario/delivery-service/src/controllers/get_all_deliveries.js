const Delivery = require("../models/Delivery");

module.exports = async (req, res) => {
  try {
    const entregas = await Delivery.find();
    res.json(entregas);
  } catch (err) {
    res.status(500).json({ error: "Error al listar entregas" });
  }
};
