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
	if conn == nil {
		panic("missing database connection")
	}
	return &PostgresMemoryDatabase{
		DB: conn,
	}
}

func (pm *PostgresMemoryDatabase) List(ctx context.Context, list_mem_id string) ([]model.Memory, error) {
	query := `SELECT * FROM users
			WHERE memory_id = $1`
	row := pm.DB.QueryRowContext(ctx, query, list_mem_id)
	m := model.Memory{}
	err := row.Scan(&m.MemoryDate, &m.MemoryPics, &m.MemoryNote)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return nil, nil
}

func (pm *PostgresMemoryDatabase) GetOne(ctx context.Context, mem_id string) (*model.Memory, error) {
	query := `SELECT * FROM users
			WHERE email = $1`
	row := pm.DB.QueryRowContext(ctx, query, mem_id)
	m := model.Memory{}
	err := row.Scan(&m.MemoryDate, &m.MemoryPics, &m.MemoryNote)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &m, nil
}

func (pm *PostgresMemoryDatabase) Hide(ctx context.Context, id string) error {
	return nil
}
