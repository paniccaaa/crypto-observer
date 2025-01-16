package v1

import (
	"fmt"

	"github.com/paniccaaa/crypto-observer/internal/model"
)

// only for coin table
func (s *Service) Create(cryptocurrency string) (model.Coin, error) {
	coin, err := s.cryptoRepository.Create(cryptocurrency)
	if err != nil {
		return model.Coin{}, fmt.Errorf("create coin: %w", err)
	}

	return coin, nil
}
