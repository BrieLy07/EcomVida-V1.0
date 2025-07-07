class Categoria:
    def __init__(self, id, nombre, descripcion):
        self.id = id
        self.nombre = nombre
        self.descripcion = descripcion
        self.productos = []  # Composición de productos

    def agregar_producto(self, producto):
        self.productos.append(producto)

    def to_dict(self, incluir_productos=False):
        base = {
            "id": self.id,
            "nombre": self.nombre,
            "descripcion": self.descripcion,
        }
        if incluir_productos:
            base["productos"] = [p.to_dict() for p in self.productos]
        return base
