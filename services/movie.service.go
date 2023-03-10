package services

import "example.com/goflix/models"

type MovieServices interface {
	CreateMovie(*models.Movie) error
	GetMovie(*string) (*models.Movie, error)
	GetAll() ([]*models.Movie, error)
	UpdateMovie(*models.Movie) error
	DeleteMovie(*string) error
}
