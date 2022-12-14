package orm

import (
	"project-backend/model"
	"project-backend/store"

	"gorm.io/gorm"
)

type ScoreRepository struct {
	conn *gorm.DB
}

func NewScoreRepository(conn *gorm.DB) store.ScoreRepository {
	return &ScoreRepository{conn: conn}
}

func (s *ScoreRepository) ScoresByUser(userID uint) ([]model.Score, error) {
	scores := []model.Score{}
	if err := s.conn.Where("user_id = ?", userID).Find(&scores).Error; err != nil {
		return nil, err
	}
	return scores, nil
}  

func (s *ScoreRepository) Create(score *model.Score) error {
	return s.conn.Create(score).Error
}