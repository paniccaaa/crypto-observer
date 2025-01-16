package api

import (
	"encoding/json"
	"net/http"

	desc "github.com/paniccaaa/crypto-observer/internal/pb"
)

func (i *Implementation) DeleteCurrencyRemove(w http.ResponseWriter, r *http.Request, params desc.DeleteCurrencyRemoveParams) {
	if err := i.cryptoService.Delete(params.Coin); err != nil {
		http.Error(w, "failed to remove cryptocurrency: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := desc.RemoveCurrencyResponse{
		Message: "Cryptocurrency removed successfully",
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode resp: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
