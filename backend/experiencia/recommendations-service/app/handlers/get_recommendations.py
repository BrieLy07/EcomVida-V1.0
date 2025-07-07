from app.grpc.generated import recommendation_pb2, recommendation_pb2_grpc
from app.chain.personalized_handler import PersonalizedRecommendationHandler

class GetRecommendationsServicer(recommendation_pb2_grpc.RecommendationServiceServicer):
    def GetRecommendations(self, request, context):
        handler = PersonalizedRecommendationHandler()
        recommendations = handler.handle(request.user_id)

        return recommendation_pb2.RecommendationList(
            items=[
                recommendation_pb2.Recommendation(
                    product_id=rec["product_id"],
                    product_name=rec["product_name"],
                    image_url=rec["image_url"]
                )
                for rec in recommendations
            ]
        )