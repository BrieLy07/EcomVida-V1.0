from app.chain.handler import Handler
from app.database.models import get_personalized_recommendations

class PersonalizedRecommendationHandler(Handler):
    def handle(self, user_id):
        recommendations = get_personalized_recommendations(user_id)
        if recommendations:
            return recommendations
        return super().handle(user_id)
