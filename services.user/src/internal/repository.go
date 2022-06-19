package user

import (
	"github.com/aliaydins/microservice/services.user/src/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db.Logger.LogMode(logger.Info)
	return &Repository{db: db}
}

func (r *Repository) Create(user *entity.User) (*entity.User, error) {
	err := r.db.Model(&entity.User{}).Create(&user).Error
	return user, err
}

func (r *Repository) FindByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
