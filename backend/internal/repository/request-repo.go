package repository

import (
	"context"

	"github.ducminhsw.prepare-project/internal/model"
)

type RequestRepositoryInterface interface {
	Get(ctx context.Context, email string) (*model.Request, error)
	Make(ctx context.Context, from string, target string, note string) error
	Break(ctx context.Context, id string) error
}
