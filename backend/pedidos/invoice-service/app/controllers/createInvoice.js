const Factura = require("../models/Factura");
const FacturaSimple = require("../strategies/FacturaSimple");
const FacturaDetallada = require("../strategies/FacturaDetallada");

module.exports = async (req, res) => {
  try {
    const { tipo, ...datos } = req.body;

    let strategy;
    if (tipo === "detallada") {
      strategy = new FacturaDetallada();
    } else {
      strategy = new FacturaSimple();
    }

    const facturaGenerada = strategy.generar(datos);
    const nuevaFactura = new Factura(facturaGenerada);
    const guardada = await nuevaFactura.save();

    res.status(201).json(guardada);
  } catch (err) {
    res.status(500).json({ error: "Error al generar factura" });
  }
};
