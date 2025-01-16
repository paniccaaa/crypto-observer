package api

import (
	"time"

	"github.com/paniccaaa/crypto-observer/internal/model"
	desc "github.com/paniccaaa/crypto-observer/internal/pb"
)

type Service interface {
	Create(cryptocurrency string) (model.Coin, error)
	Get(cryptocurrency string, timestamp time.Time) (model.CoinPrice, error)
	Delete(cryptocurrency string) error
}

var _ desc.ServerInterface = (*Implementation)(nil)

type Implementation struct {
	cryptoService Service
}

func NewImplementation(cryptoService Service) *Implementation {
	return &Implementation{
		cryptoService: cryptoService,
	}
}
