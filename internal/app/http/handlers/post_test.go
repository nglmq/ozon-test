package handlers

import (
	"github.com/nglmq/ozon-test/internal/app/service"
	"github.com/nglmq/ozon-test/internal/storage/inmemory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlePost(t *testing.T) {
	type want struct {
		code int
	}
	tests := []struct {
		name        string
		requestBody string
		request     string
		want        want
	}{
		// TODO: Add test cases.
		{
			name: "simple test 1",
			want: want{
				code: http.StatusCreated,
			},
			requestBody: `{"url":"ozon.ru"}`,
			request:     "/",
		},
		{
			name: "No URL provided test 1",
			want: want{
				code: http.StatusBadRequest,
			},
			requestBody: "",
			request:     "/",
		},
		{
			name: "simple test 2",
			want: want{
				code: http.StatusCreated,
			},
			requestBody: `{"url":"ozon.ru"}`,
			request:     "/",
		},
		{
			name: "very long url",
			want: want{
				code: http.StatusCreated,
			},
			requestBody: `{"url": "https://www.google.com/search?q=dbjksvvvvvvvvvvvvvvvvvvvvaipsjfiopqhweuofbeiwuqbfibqwib"}`,
			request:     "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := inmemory.NewInMemoryURLStorage()
			service := service.NewURLService(store)

			handler := HandlePost(service)

			request := httptest.NewRequest(http.MethodPost, tt.request, strings.NewReader(tt.requestBody))
			w := httptest.NewRecorder()

			handler(w, request)

			result := w.Result()

			assert.Equal(t, tt.want.code, result.StatusCode)

			defer result.Body.Close()
			_, err := io.ReadAll(result.Body)

			require.NoError(t, err)
		})
	}
}
