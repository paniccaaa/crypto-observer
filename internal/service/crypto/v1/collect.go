package v1

import (
	"log"
	"time"

	"github.com/paniccaaa/crypto-observer/internal/model"
	"github.com/paniccaaa/crypto-observer/internal/util"
)

func (s *Service) StartPriceUpdater() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			coins, err := s.cryptoRepository.GetAll()
			if err != nil {
				log.Printf("Error fetching coins: %v", err)
				continue
			}

			for _, coin := range coins {
				price, err := util.GetCoinPrice(coin.Name)
				if err != nil {
					log.Printf("get coin price %s: %v", coin.Name, err)
					continue
				}

				coinPrice := model.CoinPrice{
					CoinID:    coin.ID,
					Price:     int(price),
					Timestamp: time.Now(),
				}
				err = s.cryptoRepository.Save(coinPrice)
				if err != nil {
					log.Printf("save coin price %v", err)
				} else {
					log.Printf("saved price for %s: %f at %s", coin.Name, price, time.Now())
				}
			}
		}
	}
}
