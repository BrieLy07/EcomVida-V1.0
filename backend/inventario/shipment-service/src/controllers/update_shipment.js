const Mediator = require("../mediator/mediator");

module.exports = async (req, res) => {
  try {
    const { id } = req.params;
    const data = req.body;
    const result = await Mediator.notify("actualizar-envio", { id, data });
    res.json(result);
  } catch (err) {
    res.status(500).json({ error: "Error al actualizar envío" });
  }
};
