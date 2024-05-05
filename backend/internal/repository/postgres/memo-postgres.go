package postgres

import (
	"context"
	"database/sql"

	"github.ducminhsw.prepare-project/internal/model"
	"github.ducminhsw.prepare-project/internal/repository"
)

type PostgresMemoryDatabase struct {
	DB *sql.DB
}

func NewMemoryInterface(conn *sql.DB) repository.MemoRepositoryInterface {
	return &PostgresMemoryDatabase{
		DB: conn,
	}
}

func (pm *PostgresMemoryDatabase) List(ctx context.Context, id string) ([]model.Memory, error) {
	return nil, nil
}

func (pm *PostgresMemoryDatabase) GetOne(ctx context.Context) (model.Memory, error) {
	return model.Memory{}, nil
}

func (pm *PostgresMemoryDatabase) Hide(ctx context.Context, id string) error {
	return nil
}
