package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ChristianSch/gqlgen-pokemon/graph/generated"
	"github.com/ChristianSch/gqlgen-pokemon/graph/model"
)

// in-memory database
var PokemonData []*model.Pokemon
var BattleData []*model.Battle

func (r *mutationResolver) CreatePokemon(ctx context.Context, input model.NewPokemon) (*model.Pokemon, error) {
	newPokemon := model.Pokemon{
		Name:        input.Name,
		Power:       input.Power,
		Description: &input.Description,
	}

	PokemonData = append(PokemonData, &newPokemon)

	return &newPokemon, nil
}

func (r *mutationResolver) DeletePokemon(ctx context.Context, id string) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented")) // TODO:
}

func (r *mutationResolver) UpdatePokemon(ctx context.Context, input model.NewPokemon) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented")) // TODO:
}

func (r *mutationResolver) CreateBattle(ctx context.Context, input model.NewBattle) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented")) // TODO:
}

func (r *mutationResolver) DeleteBattle(ctx context.Context, id string) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented")) // TODO:
}

func (r *mutationResolver) UpdateBattle(ctx context.Context, id string) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented")) // TODDO:
}

func (r *queryResolver) GetPokemon(ctx context.Context, id string) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented")) // TODO:
}

func (r *queryResolver) ListPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	return PokemonData, nil
}

func (r *queryResolver) GetBattle(ctx context.Context, id string) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented")) // TODO:
}

func (r *queryResolver) ListBattle(ctx context.Context) ([]*model.Battle, error) {
	panic(fmt.Errorf("not implemented")) // TODO:
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
