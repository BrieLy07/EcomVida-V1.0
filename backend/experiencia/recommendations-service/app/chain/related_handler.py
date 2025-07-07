from app.chain.handler import Handler
from app.database.models import get_related_products

class RelatedRecommendationHandler(Handler):
    def handle(self, product_id):
        recommendations = get_related_products(product_id)
        if recommendations:
            return recommendations
        return super().handle(product_id)
