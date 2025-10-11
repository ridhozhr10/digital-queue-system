package repository

import (
	"context"
	"database/sql"
	"digital-queue-system/internal/model"
	"errors"

	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrDuplicateUser = errors.New("user with this username or email already exists")
)

type UserRepository struct {
	*PostgresRepository
}

func NewUserRepository(db *PostgresRepository) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByUsernameOrEmail(ctx context.Context, userIdentity string) (*model.User, error) {
	query := `SELECT id, username, password_hash FROM users WHERE username = $1 or email = $1`
	row := r.DB.QueryRowContext(ctx, query, userIdentity)

	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		log.Error().Err(err).Msg("failed to get user")
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // Unique violation
			return ErrDuplicateUser
		}
		log.Error().Err(err).Msg("failed to create user")
		return err
	}
	return nil
}
