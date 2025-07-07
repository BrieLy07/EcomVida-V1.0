class VarianteBuilder {
  constructor() {
    this.variante = {};
  }

  setProductoId(id) {
    this.variante.productoId = id;
    return this;
  }

  setColor(color) {
    this.variante.color = color;
    return this;
  }

  setTalla(talla) {
    this.variante.talla = talla;
    return this;
  }

  setPrecioExtra(precio) {
    this.variante.precioExtra = precio;
    return this;
  }

  setDisponible(disponible) {
    this.variante.disponible = disponible;
    return this;
  }

  build() {
    return this.variante;
  }
}

module.exports = VarianteBuilder;
