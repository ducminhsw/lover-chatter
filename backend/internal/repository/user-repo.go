package repository

import (
	"context"

	"github.ducminhsw.prepare-project/internal/model"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user model.User) error
	Retrieve(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, user model.User) error
}
