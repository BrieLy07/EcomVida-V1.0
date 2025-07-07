from app.templates.template_method import ProcesadorPago

class PagoTarjeta(ProcesadorPago):
    def validar(self, datos):
        assert "numero" in datos and len(datos["numero"]) == 16

    def ejecutar_pago(self, datos):
        return {"exito": True, "mensaje": "Pago aprobado"}

    def guardar_transaccion(self, datos, resultado):
        # Aquí podrías guardar logs o enviar eventos
        pass
