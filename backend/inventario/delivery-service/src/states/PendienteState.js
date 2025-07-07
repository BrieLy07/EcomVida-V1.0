class PendienteState {
  constructor(delivery) {
    this.delivery = delivery;
  }

  avanzar() {
    this.delivery.estado = "en_transito";
    return this.delivery;
  }

  cancelar() {
    this.delivery.estado = "cancelado";
    return this.delivery;
  }
}

module.exports = PendienteState;
