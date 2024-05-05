package postgres

import (
	"context"
	"database/sql"

	"github.ducminhsw.prepare-project/internal/model"
	"github.ducminhsw.prepare-project/internal/repository"
)

type PostgresUserDatabase struct {
	DB *sql.DB
}

func NewUserInterface(db *sql.DB) repository.UserRepositoryInterface {
	if db == nil {
		panic("missing database connection")
	}
	return &PostgresUserDatabase{DB: db}
}

func NewUser() *model.User {
	return &model.User{
		Email:        "",
		HashPassword: "",
		Username:     "",
		HeartKey:     "",
		LoverName:    "",
		MessagesId:   "",
	}
}

func (pu *PostgresUserDatabase) Create(ctx context.Context, user model.User) error {
	query := `INSERT INTO users 
			(email, hashpassword, username, heartkey, lovername, message_id)
			VALUES 
			($1, $2, $3, $4, $5);`
	_, err := pu.DB.ExecContext(ctx, query, user.Email, user.HashPassword, user.Username, user.LoverName, user.MessagesId)
	if err != nil {
		return err
	}
	return nil
}

func (pu *PostgresUserDatabase) Retrieve(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT * FROM users
			WHERE email = $1`
	row := pu.DB.QueryRowContext(ctx, query, email)
	u := model.User{}
	err := row.Scan(&u.Email, &u.HashPassword, &u.Username, &u.HeartKey, &u.LoverName, &u.MessagesId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &u, nil
}

func (pu *PostgresUserDatabase) Update(ctx context.Context, user model.User) error {
	uquery := `UPDATE users
		SET
			hashpassword = $1
			username = $2
			heartkey = $3
			lovername = $4
			messageid = $5
		WHERE
			email = $6`
	_, err := pu.DB.ExecContext(ctx, uquery, user.HashPassword, user.Username, user.HashPassword, user.HeartKey, user.LoverName, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (pu *PostgresUserDatabase) Delete(ctx context.Context, user model.User) error {
	return nil
}
