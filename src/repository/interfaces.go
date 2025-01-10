package repository

import (
	"coin-App/src/models"
	"time"
)

type Repository interface {
	CreateCoin(request models.Coin) error
	UpdateCoin(coinId uint, request models.Coin) error
	FindById(coinId uint) (models.Coin, error)
	FindByName(coinName string) (models.Coin, error)
	ViewAllCoins() ([]models.Coin, error)
	DeleteCoin(coinID uint) error
	LogExpiredCoins(name string, expired_date time.Time) error
	ViewExpiredCoins() ([]models.ExpiredCoinLogs, error)
}
