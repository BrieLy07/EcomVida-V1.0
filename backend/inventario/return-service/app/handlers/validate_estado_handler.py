from app.handlers.base_handler import BaseHandler

class ValidateEstadoHandler(BaseHandler):
    def handle(self, return_data):
        estado = return_data.get("estado")
        if estado and estado not in ["pendiente", "aprobado", "rechazado"]:
            raise ValueError("Estado inválido")
        return super().handle(return_data)
