package services

import (
	"context"
	"errors"

	"example.com/goflix/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieServicesImpl struct {
	moviecollection *mongo.Collection
	ctx             context.Context
}

func NewMovieService(moviecollection *mongo.Collection, ctx context.Context) MovieServices {
	return &MovieServicesImpl{
		moviecollection: moviecollection,
		ctx:             ctx,
	}
}

func (u *MovieServicesImpl) CreateMovie(movie *models.Movie) error {
	_, err := u.moviecollection.InsertOne(u.ctx, movie)
	return err
}

func (u *MovieServicesImpl) GetMovie(name *string) (*models.Movie, error) {
	var movie *models.Movie
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.moviecollection.FindOne(u.ctx, query).Decode(&movie)

	return movie, err
}

func (u *MovieServicesImpl) GetAll() ([]*models.Movie, error) {
	var movies []*models.Movie
	cursor, err := u.moviecollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var movie models.Movie
		err := cursor.Decode(&movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(movies) == 0 {
		return nil, errors.New("documents not found")
	}
	return movies, nil
}

func (u *MovieServicesImpl) UpdateMovie(movie *models.Movie) error {
	filter := bson.D{primitive.E{Key: "name", Value: movie.Name}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: movie.Name}, primitive.E{Key: "age", Value: movie.Duration}, primitive.E{Key: "address", Value: movie.Link}, primitive.E{Key: "age", Value: movie.Publisher}}}}
	result, _ := u.moviecollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *MovieServicesImpl) DeleteMovie(name *string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	result, _ := u.moviecollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
