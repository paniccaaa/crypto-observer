package v1

import (
	"time"

	"github.com/paniccaaa/crypto-observer/internal/model"
)

type Storage interface {
	Create(cryptocurrency string) (model.Coin, error)
	Get(cryptocurrency string, timestamp time.Time) (model.CoinPrice, error)
	Delete(cryptocurrency string) error
	GetAll() ([]model.Coin, error)
	Save(coinPrice model.CoinPrice) error
}

type Service struct {
	cryptoRepository Storage
}

func NewService(cryptoRepo Storage) *Service {
	return &Service{
		cryptoRepository: cryptoRepo,
	}
}
