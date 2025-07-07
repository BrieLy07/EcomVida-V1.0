from concurrent import futures
import grpc
import os
from dotenv import load_dotenv

from app.grpc.generated import recommendation_pb2_grpc
from app.handlers.get_recommendations import GetRecommendationsServicer
from app.handlers.get_related import GetRelatedProductsServicer
from app.handlers.post_feedback import PostFeedbackServicer

load_dotenv()

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    
    recommendation_pb2_grpc.add_RecommendationServiceServicer_to_server(
        GetRecommendationsServicer(), server
    )
    recommendation_pb2_grpc.add_RecommendationServiceServicer_to_server(
        GetRelatedProductsServicer(), server
    )
    recommendation_pb2_grpc.add_RecommendationServiceServicer_to_server(
        PostFeedbackServicer(), server
    )

    server.add_insecure_port(f"[::]:{os.getenv('APP_PORT')}")
    print(f"gRPC server running on port {os.getenv('APP_PORT')}")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()
