const Mediator = require("../mediator/mediator");

module.exports = async (req, res) => {
  try {
    const result = await Mediator.notify("obtener-envio", req.params.id);
    res.json(result);
  } catch (err) {
    res.status(404).json({ error: "Envío no encontrado" });
  }
};
