package v1

import "fmt"

func (s *Service) Delete(cryptocurrency string) error {
	if err := s.cryptoRepository.Delete(cryptocurrency); err != nil {
		return fmt.Errorf("delete cryptocurrency: %w", err)
	}

	return nil
}
