package orm

import (
	"fmt"
	"project-backend/config"
	"project-backend/model"
	"project-backend/store"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB

	users  store.UserRepository
	scores store.ScoreRepository
}

func NewDB() store.Store {
	return &DB{}
}

func (db *DB) UserRepository() store.UserRepository {
	if db.users == nil {
		db.users = NewUserRepository(db.conn)
	}
	return db.users
}

func (db *DB) ScoreRepository() store.ScoreRepository {
	if db.scores == nil {
		db.scores = NewScoreRepository(db.conn)
	}

	return db.scores
}

func (db *DB) Connect(config *config.Config) error {
	address := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBName,
		config.DBPassword,
	)

	conn, err := gorm.Open(postgres.Open(address))
	if err != nil {
		return err
	}

	if err := conn.AutoMigrate(
		&model.User{}, 
		&model.Score{},
	); err != nil {
		return err
	}

	db.conn = conn
	return nil
}
