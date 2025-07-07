from functools import wraps

def log_operacion(mensaje: str):
    def decorador(func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            print(f"[LOG] {mensaje}")
            return func(*args, **kwargs)
        return wrapper
    return decorador
