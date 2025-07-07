from app.grpc.generated import recommendation_pb2, recommendation_pb2_grpc
from app.database.models import save_feedback

class PostFeedbackServicer(recommendation_pb2_grpc.RecommendationServiceServicer):
    def SendFeedback(self, request, context):
        save_feedback(
            user_id=request.user_id,
            product_id=request.product_id,
            feedback_type=request.feedback_type
        )
        return recommendation_pb2.FeedbackResponse(message="Feedback recibido correctamente.")
















