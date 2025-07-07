class BaseHandler:
    def __init__(self):
        self._next = None

    def set_next(self, handler):
        self._next = handler
        return handler

    def handle(self, return_data):
        if self._next:
            return self._next.handle(return_data)
        return return_data
