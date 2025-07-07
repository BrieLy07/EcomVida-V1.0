const PendienteState = require("./PendienteState");
const EnTransitoState = require("./EnTransitoState");

function getStateInstance(delivery) {
  switch (delivery.estado) {
    case "pendiente":
      return new PendienteState(delivery);
    case "en_transito":
      return new EnTransitoState(delivery);
    default:
      return null;
  }
}

module.exports = getStateInstance;
