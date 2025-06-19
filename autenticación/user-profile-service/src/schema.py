import graphene
from resolvers import Query, Mutation

schema = graphene.Schema(query=Query, mutation=Mutation)
