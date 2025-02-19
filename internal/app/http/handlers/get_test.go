package handlers

import (
	"context"
	"github.com/nglmq/ozon-test/internal/app/service"
	"github.com/nglmq/ozon-test/internal/storage/inmemory"
	"github.com/nglmq/ozon-test/pkg/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGet(t *testing.T) {
	storage := inmemory.NewInMemoryURLStorage()
	urlService := service.NewURLService(storage)

	storage.Save(context.Background(), &models.URL{Original: "https://ozon.ru", Short: "abcdef"})

	handler := HandleGet(urlService)

	tests := []struct {
		name       string
		urlPath    string
		statusCode int
		location   string
	}{
		{
			name:       "Valid URL",
			urlPath:    "/abcdef",
			statusCode: http.StatusTemporaryRedirect,
			location:   "https://ozon.ru",
		},
		{
			name:       "Invalid URL",
			urlPath:    "/nonexistent",
			statusCode: http.StatusBadRequest,
			location:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.urlPath, nil)
			w := httptest.NewRecorder()

			handler(w, req)

			resp := w.Result()

			assert.Equal(t, tt.statusCode, resp.StatusCode)

			if tt.statusCode == http.StatusTemporaryRedirect {
				assert.Equal(t, tt.location, resp.Header.Get("Location"))
			}
		})
	}
}
