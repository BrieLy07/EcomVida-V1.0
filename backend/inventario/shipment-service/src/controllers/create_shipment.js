const Mediator = require("../mediator/mediator");

module.exports = async (req, res) => {
  try {
    const result = await Mediator.notify("crear-envio", req.body);
    res.status(201).json(result);
  } catch (err) {
    res.status(500).json({ error: "Error al crear envío" });
  }
};
