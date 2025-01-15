package pg

import (
	"log"

	"github.com/jmoiron/sqlx"
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
