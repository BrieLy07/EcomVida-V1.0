const Mediator = {
  handlers: {},

  register(event, handler) {
    this.handlers[event] = handler;
  },

  async notify(event, payload) {
    const handler = this.handlers[event];
    if (!handler) throw new Error(`No handler for event: ${event}`);
    return await handler(payload);
  }
};

// Registro de eventos
Mediator.register("crear-envio", require("../services/createShipment"));
Mediator.register("actualizar-envio", async ({ id, data }) =>
  require("../services/updateShipment")(id, data)
);
Mediator.register("eliminar-envio", require("../services/deleteShipment"));
Mediator.register("obtener-envio", require("../services/getShipmentById"));
Mediator.register("listar-envios", require("../services/getAllShipments"));

module.exports = Mediator;
