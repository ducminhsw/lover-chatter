package repository

import (
	"context"

	"github.ducminhsw.prepare-project/internal/model"
)

type MemoRepositoryInterface interface {
	List(ctx context.Context, list_mem_id string) ([]model.Memory, error)
	GetOne(ctx context.Context, mem_id string) (*model.Memory, error)
	Hide(ctx context.Context, id string) error
}
