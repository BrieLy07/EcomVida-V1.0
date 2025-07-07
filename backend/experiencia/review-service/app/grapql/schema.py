import strawberry
from app.graphql.queries.review_queries import QueryReview
from app.graphql.mutations.review_mutations import MutationReview

@strawberry.type
class Query(QueryReview):
    pass

@strawberry.type
class Mutation(MutationReview):
    pass

schema = strawberry.Schema(query=Query, mutation=Mutation)
