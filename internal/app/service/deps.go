package service

import (
	"context"
	"github.com/nglmq/ozon-test/pkg/models"
)

type URLRepository interface {
	Save(ctx context.Context, url *models.URL) error
	GetOriginal(ctx context.Context, short string) (*models.URL, error)
	GetShort(ctx context.Context, original string) (*models.URL, error)
}
