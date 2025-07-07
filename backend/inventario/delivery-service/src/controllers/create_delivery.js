const Delivery = require("../models/Delivery");

module.exports = async (req, res) => {
  try {
    const nuevaEntrega = new Delivery(req.body);
    const guardada = await nuevaEntrega.save();
    res.status(201).json(guardada);
  } catch (err) {
    res.status(400).json({ error: "Error al registrar entrega" });
  }
};
