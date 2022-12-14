package orm

import (
	"project-backend/model"
	"project-backend/store"

	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) store.UserRepository {
	return &UserRepository{conn: conn}
}

func (u *UserRepository) UserByUsername(username string) (*model.User, error) {
	user := model.User{}
	if err := u.conn.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UserByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := u.conn.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UserById(id uint) (*model.User, error) {
	user := model.User{}
	if err := u.conn.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Create(user *model.User) error {
	return u.conn.Create(user).Error
}