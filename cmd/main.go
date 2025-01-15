package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/paniccaaa/crypto-observer/internal/api/crypto"
	"github.com/paniccaaa/crypto-observer/internal/app"
	"github.com/paniccaaa/crypto-observer/internal/repository/crypto/pg"
	v1 "github.com/paniccaaa/crypto-observer/internal/service/crypto/v1"
)

func main() {
	cfg := app.NewConfig()

	db := pg.NewRepository(cfg.DB_URI)
	cryptoService := v1.NewService(db)

	impl := api.NewImplementation(cryptoService)

	srv := app.SetupServer(cfg, impl)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)

	<-done

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to stop server: %v", err)
	}

	log.Print("server stopped")
}
