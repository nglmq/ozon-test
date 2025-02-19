package service

import (
	"context"
	"errors"
	"github.com/nglmq/ozon-test/internal/config"
	"github.com/nglmq/ozon-test/internal/storage"
	"github.com/nglmq/ozon-test/pkg/models"
	"github.com/nglmq/ozon-test/pkg/shorten"
)

type URLService struct {
	repository URLRepository
}

func NewURLService(repository URLRepository) *URLService {
	return &URLService{repository: repository}
}

func (s *URLService) ShortenURL(ctx context.Context, original string) (*models.URLResponse, error) {
	short := shorten.NewRandomURL()
	url := &models.URL{
		Original: original,
		Short:    short,
	}

	if err := s.repository.Save(ctx, url); err != nil {
		if errors.Is(err, storage.ErrURLExists) {
			oldShort, err := s.repository.GetShort(ctx, original)
			if err != nil {
				return nil, err
			}

			return &models.URLResponse{
				Short: config.BaseURL + "/" + oldShort.Short,
			}, nil
		}
		return nil, err
	}

	fullURL := config.BaseURL + "/" + short

	return &models.URLResponse{
		Short: fullURL,
	}, nil
}

func (s *URLService) GetOriginalURL(ctx context.Context, short string) (*models.URL, error) {
	return s.repository.GetOriginal(ctx, short)
}

func (s *URLService) GetShortURL(ctx context.Context, original string) (*models.URL, error) {
	return s.repository.GetShort(ctx, original)
}
