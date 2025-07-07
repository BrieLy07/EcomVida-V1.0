class Handler:
    def __init__(self):
        self.next = None

    def set_next(self, handler):
        self.next = handler
        return handler

    def handle(self, request_data):
        if self.next:
            return self.next.handle(request_data)
        return []