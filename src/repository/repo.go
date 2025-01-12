package repository

import (
	"coin-App/src/models"
	"time"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) CreateCoin(request models.Coin) error {
	if err := r.db.Create(&request).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) UpdateCoin(coinId uint, request models.Coin) error {
	result := r.db.Model(&models.Coin{}).Where("id = ?", coinId).Updates(request)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *repo) FindById(coinID uint) (models.Coin, error) {
	var coin models.Coin
	if err := r.db.First(&coin, coinID).Error; err != nil {
		return models.Coin{}, err
	}
	return coin, nil
}

func (r *repo) FindByName(coinName string) (models.Coin, error) {
	var coin models.Coin
	if err := r.db.Where("name = ?", coinName).First(&coin).Error; err != nil {
		return models.Coin{}, err
	}
	return coin, nil
}

func (r *repo) ViewAllCoins() ([]models.Coin, error) {
	var coins []models.Coin
	if err := r.db.Find(&coins).Error; err != nil {
		return nil, err
	}
	return coins, nil
}

func (r *repo) DeleteCoin(coinID uint) error {
	if err := r.db.Delete(&models.Coin{}, coinID).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) LogExpiredCoins(name string, expiredDate time.Time) error {
	log := models.ExpiredCoinLogs{
		Name:      name,
		ExpiredAt: expiredDate,
	}
	if err := r.db.Create(&log).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) ViewExpiredCoins() ([]models.ExpiredCoinLogs, error) {
	var expiredCoinLogs []models.ExpiredCoinLogs
	if err := r.db.Find(&expiredCoinLogs).Error; err != nil {
		return nil, err
	}
	return expiredCoinLogs, nil
}

func (r *repo) CheckByName(name string) (models.ExpiredCoinLogs, error) {
	var coin models.ExpiredCoinLogs
	if err := r.db.Where("name = ?", name).First(&coin).Error; err != nil {
		return models.ExpiredCoinLogs{}, err
	}
	return coin, nil
}
