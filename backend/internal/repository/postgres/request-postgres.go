package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.ducminhsw.prepare-project/internal/model"
	"github.ducminhsw.prepare-project/internal/repository"
)

type PostgresRequestDatabase struct {
	DB *sql.DB
}

func NewFriendDatabase(conn *sql.DB) repository.RequestRepositoryInterface {
	return &PostgresRequestDatabase{DB: conn}
}

func (pf *PostgresRequestDatabase) Get(ctx context.Context, email string) (*model.Request, error) {
	query := `SELECT * FROM requests WHERE target = $1`
	res := pf.DB.QueryRowContext(ctx, query, email)

	req := model.Request{}

	err := res.Scan(&req.Id, &req.From, &req.Target, &req.Note)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &req, nil
}

func (pf *PostgresRequestDatabase) Make(ctx context.Context, from, target, note string) error {
	query := `INSERT INTO requests (id, from, target, note) VALUES ($1, $2, $3, $4)`
	time_id := time.Now().String()
	_, err := pf.DB.ExecContext(ctx, query, time_id, from, target, note)
	if err != nil {
		return err
	}

	return nil
}

func (pf *PostgresRequestDatabase) Break(ctx context.Context, id string) error {
	query := `SELECT * FROM requests WHERE id = $1`
	_, err := pf.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	query = `DELETE FROM requests WHERE id = $1`
	_, err = pf.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
