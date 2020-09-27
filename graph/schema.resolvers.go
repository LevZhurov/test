package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"test1/graph/generated"
	"test1/graph/model"
)

func (r *messageResolver) Timestamp(ctx context.Context, obj *model.Message) (string, error) {
	panic(fmt.Errorf("not implemented 1"))
}

func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	panic(fmt.Errorf("not implemented 2"))
}

func (r *queryResolver) Messages(ctx context.Context, limit *int, offset *int) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented 3"))
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type messageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
