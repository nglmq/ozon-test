package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgerrcode"
	_ "github.com/jackc/pgx/v5"
	"github.com/lib/pq"
	"github.com/nglmq/ozon-test/internal/storage"
	"github.com/nglmq/ozon-test/pkg/models"
)

type PostgresURLStorage struct {
	db *sql.DB
}

func NewPostgresURLStorage(dsn string) *PostgresURLStorage {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls(
 		id SERIAL PRIMARY KEY,
		original TEXT NOT NULL,
		short TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP);
	`)
	if err != nil {
		return nil
	}

	return &PostgresURLStorage{
		db: db,
	}
}

func (s *PostgresURLStorage) Save(ctx context.Context, url *models.URL) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := s.db.ExecContext(ctx,
		"INSERT INTO urls (original, short) VALUES ($1, $2)",
		url.Original,
		url.Short)
	if err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code == pgerrcode.UniqueViolation {
			return storage.ErrURLExists
		}
		return fmt.Errorf("failed to save URL: %w", err)
	}

	return nil
}

func (s *PostgresURLStorage) GetOriginal(ctx context.Context, short string) (*models.URL, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var url models.URL

	err := s.db.QueryRowContext(ctx,
		"SELECT (original, short) FROM urls WHERE short = $1", short).
		Scan(&url.Original, &url.Short)
	if err != nil {
		return nil, storage.ErrURLNotFound
	}

	return &url, nil
}

func (s *PostgresURLStorage) GetShort(ctx context.Context, original string) (*models.URL, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var url models.URL

	err := s.db.QueryRowContext(ctx,
		"SELECT (original, short) FROM urls WHERE original = $1", original).
		Scan(&url.Original, &url.Short)
	if err != nil {
		return nil, storage.ErrURLNotFound
	}

	return &url, nil
}
