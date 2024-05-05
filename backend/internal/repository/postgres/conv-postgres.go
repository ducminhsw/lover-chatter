package postgres

import (
	"context"
	"database/sql"

	"github.ducminhsw.prepare-project/internal/model"
	"github.ducminhsw.prepare-project/internal/repository"
)

type PostgresConversationDatabase struct {
	DB *sql.DB
}

func NewConversationInterface(conn *sql.DB) repository.ConversationRepositoryInterface {
	return &PostgresConversationDatabase{
		DB: conn,
	}
}

func (pc *PostgresConversationDatabase) Create(ctx context.Context, id string) error {
	return nil
}

func (pc *PostgresConversationDatabase) Retrieve(ctx context.Context, id string) (*model.Conversation, error) {
	return nil, nil
}

func (pc *PostgresConversationDatabase) Update(ctx context.Context, id string) error {
	return nil
}

func (pc *PostgresConversationDatabase) Delete(ctx context.Context, id string) error {
	return nil
}
