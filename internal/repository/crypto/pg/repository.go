package pg

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/paniccaaa/crypto-observer/internal/model"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(DB_URI string) *Repository {
	db, err := sqlx.Connect("postgres", DB_URI)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to verify connection to db: %v", err)
	}

	return &Repository{db: db}
}

func (r *Repository) Save() {

}

// coin table
func (r *Repository) Create(cryptocurrency string) (model.Coin, error) {
	query := "INSERT INTO coin (name) VALUES ($1) RETURNING id, name, created_at"

	var coin model.Coin
	if err := r.db.QueryRowx(query, cryptocurrency).StructScan(&coin); err != nil {
		log.Printf("failed to create crypto: %v", err)
		return model.Coin{}, fmt.Errorf("failed to add cryptocurrency: %w", err)
	}

	return coin, nil
}

func (r *Repository) Get(cryptocurrency string, timestamp time.Time) (model.CoinPrice, error) {
	query := `
		SELECT c.id, c.name, cp.price, cp.timestamp
		FROM coin c
			INNER JOIN coin_prices cp ON c.id = cp.coin_id
		WHERE c.name = $1 AND cp.timestamp <= $2
		ORDER BY cp.timestamp DESC
		LIMIT 1`

	var coinPrice model.CoinPrice
	if err := r.db.QueryRowx(query, cryptocurrency, timestamp.Unix()).StructScan(&coinPrice); err != nil {
		log.Printf("failed to get price cryptocurrency: %v", err)
		return model.CoinPrice{}, fmt.Errorf("failed to get price: %w", err)
	}

	return coinPrice, nil
}

func (r *Repository) Delete(cryptocurrency string) error {
	query := "DELETE FROM coin WHERE name = $1"

	_, err := r.db.Exec(query, cryptocurrency)
	if err != nil {
		return fmt.Errorf("failed to delete cryptocurrency: %w", err)
	}

	return nil
}
