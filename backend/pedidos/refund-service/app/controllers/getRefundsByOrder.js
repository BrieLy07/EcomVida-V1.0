const ReembolsosPorOrden = require("../commands/ReembolsosPorOrden");

module.exports = async (req, res) => {
  try {
    const comando = new ReembolsosPorOrden(req.params.orderId);
    const resultado = await comando.ejecutar();
    res.json(resultado);
  } catch (err) {
    res.status(500).json({ error: "Error al listar reembolsos" });
  }
};
