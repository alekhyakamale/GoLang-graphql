package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/alekhyakamale/go-api/graph/generated"
	"github.com/alekhyakamale/go-api/graph/model"
	"github.com/alekhyakamale/go-api/internal/dogs"
)

func (r *mutationResolver) AddDog(ctx context.Context, input model.NewDog) (*model.Dog, error) {
	var dog dogs.Dog      //Dog struct from internal/dogs/dogs.go
	dog.Name = input.Name //accessing the dog struct
	dog.IsGoodBoi = input.IsGoodBoi
	dogID := dog.Save() //Creates an id for our dog
	return &model.Dog{ID: strconv.FormatInt(dogID, 10), Name: dog.Name, IsGoodBoi: dog.IsGoodBoi}, nil
}

func (r *mutationResolver) UpgradeDog(ctx context.Context, input model.NewDog) (*model.Dog, error) {
	var dog dogs.Dog
	dog.Name = input.Name
	dog.IsGoodBoi = input.IsGoodBoi
	dog.UpgradeDog()
	return &model.Dog{Name: dog.Name, IsGoodBoi: dog.IsGoodBoi}, nil
}

func (r *mutationResolver) UpForAdoption(ctx context.Context, input model.DogID) ([]*model.Dog, error) {
	var dog dogs.Dog
	var resultDogs []*model.Dog
	var dbDogs []dogs.Dog
	dog.ID = input.ID
	dog.UpForAdoption()
	dbDogs = dogs.GetAll()
	for _, dog := range dbDogs {
		resultDogs = append(resultDogs, &model.Dog{ID: dog.ID, Name: dog.Name, IsGoodBoi: dog.IsGoodBoi})
	}
	return resultDogs, nil
}

func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	var resultDogs []*model.Dog
	var dbDogs []dogs.Dog
	dbDogs = dogs.GetAll()
	for _, dog := range dbDogs {
		resultDogs = append(resultDogs, &model.Dog{ID: dog.ID, Name: dog.Name, IsGoodBoi: dog.IsGoodBoi})
	}
	return resultDogs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
