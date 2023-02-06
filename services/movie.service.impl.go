package services

import (
	"example.com/goflix/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieServicesImpl struct {
	moviecollection *mongo.Collection
	ctx             contxt.Context
}

func (u *MovieServicesImpl) CreateMovie(movie *models.Movie) error {
	return nil
}

func (u *MovieServicesImpl) GetMovie(name *string) (*models.Movie, error) {
	return nil, nil
}

func (u *MovieServicesImpl) GetAll() []*models.Movie {
	return nil
}

func (u *MovieServicesImpl) UpdateMovie(movie *models.Movie) error {
	return nil
}
func (u *MovieServicesImpl) DeleteMovie(name *string) error {
	return nil
}
