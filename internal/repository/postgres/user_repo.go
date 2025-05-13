package postgres

import (
	domain "splunk_soar_clone/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) CreateToken(token *domain.Token) error {
	return r.db.Create(token).Error
}

func (r *UserRepository) DeleteTokenByUserID(userID int64) error {
	return r.db.Where("user_id = ?", userID).Delete(&domain.Token{}).Error
}
