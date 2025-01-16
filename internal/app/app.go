package app

import (
	"net/http"
	"strconv"

	api "github.com/paniccaaa/crypto-observer/internal/api/crypto"
	desc "github.com/paniccaaa/crypto-observer/internal/pb"
)

func SetupServer(cfg *Config, impl *api.Implementation) *http.Server {
	mux := http.NewServeMux()

	mux.Handle("GET /currency/price", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		coin := r.FormValue("coin")
		time := r.FormValue("timestamp")

		timestamp, err := strconv.Atoi(time)
		if err != nil {
			http.Error(w, "invalid timestamp", http.StatusBadRequest)
			return
		}

		params := desc.GetCurrencyPriceParams{
			Coin:      coin,
			Timestamp: timestamp,
		}

		impl.GetCurrencyPrice(w, r, params)
	}))

	mux.Handle("POST /currency/add", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		impl.PostCurrencyAdd(w, r)
	}))

	mux.Handle("DELETE /currency/remove", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		coin := r.FormValue("coin")
		params := desc.DeleteCurrencyRemoveParams{
			Coin: coin,
		}

		impl.DeleteCurrencyRemove(w, r, params)
	}))

	srv := &http.Server{
		Addr:    cfg.Server.Addr,
		Handler: mux,
	}

	return srv
}
