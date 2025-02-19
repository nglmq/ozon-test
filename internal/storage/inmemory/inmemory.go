package inmemory

import (
	"context"
	"github.com/nglmq/ozon-test/internal/storage"
	"github.com/nglmq/ozon-test/pkg/models"
	"sync"
)

type MemoryURLStorage struct {
	store map[string]string
	rw    sync.RWMutex
}

func NewInMemoryURLStorage() *MemoryURLStorage {
	return &MemoryURLStorage{
		store: make(map[string]string),
	}
}

func (s *MemoryURLStorage) Save(ctx context.Context, url *models.URL) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.rw.Lock()
	defer s.rw.Unlock()

	for _, v := range s.store {
		if v == url.Original {
			return storage.ErrURLExists
		}
	}

	s.store[url.Short] = url.Original
	return nil
}

func (s *MemoryURLStorage) GetOriginal(ctx context.Context, short string) (*models.URL, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	s.rw.RLock()
	defer s.rw.RUnlock()

	url, exists := s.store[short]
	if !exists {
		return nil, storage.ErrURLNotFound
	}

	return &models.URL{Short: short, Original: url}, nil
}

func (s *MemoryURLStorage) GetShort(ctx context.Context, original string) (*models.URL, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	s.rw.RLock()
	defer s.rw.RUnlock()

	for short, v := range s.store {
		if v == original {
			return &models.URL{Short: short, Original: original}, nil
		}
	}

	return nil, storage.ErrURLNotFound
}
