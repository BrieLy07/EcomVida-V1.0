from app.grpc.generated import recommendation_pb2, recommendation_pb2_grpc
from app.chain.related_handler import RelatedRecommendationHandler

class GetRelatedProductsServicer(recommendation_pb2_grpc.RecommendationServiceServicer):
    def GetRelatedProducts(self, request, context):
        handler = RelatedRecommendationHandler()
        recommendations = handler.handle(request.product_id)

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