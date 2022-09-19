package history

import (
	"github.com/aliaydins/microservice/service.history/src/entity"
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

func (r *Repository) Create(history *entity.History) error {
	err := r.db.Model(&entity.History{}).Create(&history).Error
	return err
}

func (r *Repository) GetListById(id int) ([]entity.History, error) {
	var historyItems []entity.History
	err := r.db.Where(&entity.History{UserId: id}).Find(&historyItems).Error
	return historyItems, err
}
