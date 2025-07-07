const FacturaStrategy = require("./FacturaStrategy");

class FacturaSimple extends FacturaStrategy {
  generar(datos) {
    return {
      orderId: datos.orderId,
      cliente: datos.cliente,
      monto: datos.monto,
      tipo: "simple"
    };
  }
}

module.exports = FacturaSimple;
