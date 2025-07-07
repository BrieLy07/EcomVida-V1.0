from abc import ABC, abstractmethod

class ProcesadorPago(ABC):
    def procesar(self, datos):
        self.validar(datos)
        resultado = self.ejecutar_pago(datos)
        self.guardar_transaccion(datos, resultado)
        return resultado

    @abstractmethod
    def validar(self, datos):
        pass

    @abstractmethod
    def ejecutar_pago(self, datos):
        pass

    @abstractmethod
    def guardar_transaccion(self, datos, resultado):
        pass
