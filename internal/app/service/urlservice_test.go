package service

import (
	"context"
	"errors"
	"github.com/nglmq/ozon-test/pkg/models"
	"reflect"
	"testing"
)

type MockURLRepository struct {
	urls map[string]*models.URL
}

func (r *MockURLRepository) Save(ctx context.Context, url *models.URL) error {
	if url.Original == "" || url.Short == "" {
		return errors.New("invalid URL data")
	}
	r.urls[url.Short] = url
	return nil
}

func (r *MockURLRepository) GetOriginal(ctx context.Context, short string) (*models.URL, error) {
	url, exists := r.urls[short]
	if !exists {
		return nil, errors.New("URL not found")
	}
	return url, nil
}

func (r *MockURLRepository) GetShort(ctx context.Context, original string) (*models.URL, error) {
	url, exists := r.urls[original]
	if !exists {
		return nil, errors.New("URL not found")
	}
	return url, nil
}

func TestURLService_GetOriginalURL(t *testing.T) {
	mockRepo := &MockURLRepository{
		urls: map[string]*models.URL{
			"shortenedOzonURL": {
				Original: "https://ozon.ru",
				Short:    "shortenedOzonURL",
			},
		},
	}
	service := NewURLService(mockRepo)

	tests := []struct {
		name    string
		short   string
		want    *models.URL
		wantErr bool
	}{
		{
			name:  "Success - valid short URL",
			short: "shortenedOzonURL",
			want: &models.URL{
				Original: "https://ozon.ru",
				Short:    "shortenedOzonURL",
			},
			wantErr: false,
		},
		{
			name:    "Error - short URL not found",
			short:   "unknownShortURL",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetOriginalURL(context.Background(), tt.short)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOriginalURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOriginalURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
