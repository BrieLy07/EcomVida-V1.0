const Mediator = require("../mediator/mediator");

module.exports = async (req, res) => {
  try {
    const result = await Mediator.notify("eliminar-envio", req.params.id);
    res.json(result);
  } catch (err) {
    res.status(500).json({ error: "Error al eliminar envío" });
  }
};
