package repository

import (
	"context"

	"github.ducminhsw.prepare-project/internal/model"
)

type ConversationRepositoryInterface interface {
	Create(ctx context.Context, id string) error
	Retrieve(ctx context.Context, id string) (*model.Conversation, error)
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}
