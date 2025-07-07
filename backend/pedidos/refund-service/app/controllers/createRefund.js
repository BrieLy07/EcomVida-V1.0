const CrearReembolso = require("../commands/CrearReembolso");

module.exports = async (req, res) => {
  try {
    const comando = new CrearReembolso(req.body);
    const resultado = await comando.ejecutar();
    res.status(201).json(resultado);
  } catch (err) {
    res.status(500).json({ error: "Error al crear reembolso" });
  }
};
