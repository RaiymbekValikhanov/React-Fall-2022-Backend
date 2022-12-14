package store

import (
	"project-backend/config"
	"project-backend/model"
)

type Store interface {
	Connect(cfg *config.Config) error

	UserRepository() UserRepository
	ScoreRepository() ScoreRepository
}

type UserRepository interface {
	UserByUsername(username string) (*model.User, error)
	UserByEmail(email string) (*model.User, error)
	UserById(id uint) (*model.User, error)
	Create(user *model.User) error
}

type ScoreRepository interface {
	ScoresByUser(userID uint) ([]model.Score, error)
	Create(score *model.Score) error
}
