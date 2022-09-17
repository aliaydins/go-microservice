package wallet

import (
	"github.com/aliaydins/microservice/service.wallet/src/entity"
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

func (r *Repository) FindById(id int) (*entity.Wallet, error) {
	wallet := new(entity.Wallet)
	err := r.db.Where("user_id = ?", id).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *Repository) Create(wallet *entity.Wallet) error {
	err := r.db.Model(&entity.Wallet{}).Create(&wallet).Error
	if err != nil {
		return err
	}
	return nil
}
