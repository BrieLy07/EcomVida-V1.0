from app.models.producto import Producto

class ProductAdapter:

    @staticmethod
    def transformar_fila(fila):
        return Producto(*fila)
