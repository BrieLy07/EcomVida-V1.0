from app.handlers.base_handler import BaseHandler

class ValidateMotivoHandler(BaseHandler):
    def handle(self, return_data):
        if not return_data.get("motivo"):
            raise ValueError("El motivo es obligatorio")
        return super().handle(return_data)
