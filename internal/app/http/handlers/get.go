package handlers

import (
	"github.com/nglmq/ozon-test/internal/app/service"
	"net/http"
	"strings"
)

func HandleGet(service *service.URLService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET requests are allowed!", http.StatusBadRequest)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/")
		if id == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		url, err := service.GetOriginalURL(r.Context(), id)
		if err != nil {
			http.Error(w, "URL not found", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(url.Original, "http://") && !strings.HasPrefix(url.Original, "https://") {
			url.Original = "http://" + url.Original
		}

		w.Header().Set("Location", url.Original)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}
