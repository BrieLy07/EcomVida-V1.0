const FacturaStrategy = require("./FacturaStrategy");

class FacturaDetallada extends FacturaStrategy {
  generar(datos) {
    return {
      orderId: datos.orderId,
      cliente: datos.cliente,
      monto: datos.monto,
      detalle: datos.detalle,
      tipo: "detallada"
    };
  }
}

module.exports = FacturaDetallada;
