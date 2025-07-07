class EnTransitoState {
  constructor(delivery) {
    this.delivery = delivery;
  }

  avanzar() {
    this.delivery.estado = "entregado";
    return this.delivery;
  }

  cancelar() {
    this.delivery.estado = "cancelado";
    return this.delivery;
  }
}

module.exports = EnTransitoState;
