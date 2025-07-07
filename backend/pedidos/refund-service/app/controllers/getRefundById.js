const ObtenerReembolso = require("../commands/ObtenerReembolso");

module.exports = async (req, res) => {
  try {
    const comando = new ObtenerReembolso(req.params.id);
    const resultado = await comando.ejecutar();
    if (!resultado) return res.status(404).json({ error: "No encontrado" });
    res.json(resultado);
  } catch (err) {
    res.status(500).json({ error: "Error al obtener reembolso" });
  }
};
