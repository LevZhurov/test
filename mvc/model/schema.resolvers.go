package model

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"test1/graph/generated"
	"test1/graph/model"
)

func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	panic(fmt.Errorf("mutationResolver CreateMessage not implemented"))
}

func (r *queryResolver) Messages(ctx context.Context, limit *int, offset *int) ([]*model.Message, error) {
	panic(fmt.Errorf("queryResolver Messages not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
