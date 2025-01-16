package v1

import (
	"fmt"
	"time"

	"github.com/paniccaaa/crypto-observer/internal/model"
)

func (s *Service) Get(cryptocurrency string, timestamp time.Time) (model.CoinPrice, error) {
	coinPrice, err := s.cryptoRepository.Get(cryptocurrency, timestamp)
	if err != nil {
		return model.CoinPrice{}, fmt.Errorf("get cryptocurrency: %w", err)
	}

	return coinPrice, nil
}
