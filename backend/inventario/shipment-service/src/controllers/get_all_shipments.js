const Mediator = require("../mediator/mediator");

module.exports = async (req, res) => {
  try {
    const result = await Mediator.notify("listar-envios");
    res.json(result);
  } catch (err) {
    res.status(500).json({ error: "Error al listar envíos" });
  }
};
