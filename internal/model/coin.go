package model

import "time"

type Coin struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CoinPrice struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CoinID    int       `json:"coin_id"`
	Price     int       `json:"price" db:"price"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}
