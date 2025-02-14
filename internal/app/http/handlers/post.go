package handlers

import (
	"encoding/json"
	"github.com/nglmq/ozon-test/internal/app/service"
	"github.com/nglmq/ozon-test/pkg/models"
	"net/http"
)

func HandlePost(service *service.URLService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.URLRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		urlResponse, err := service.ShortenURL(r.Context(), req.URL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(urlResponse)
	}
}
