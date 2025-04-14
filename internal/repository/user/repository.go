package user_repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nikita-reshetnyak/auth/internal/client/db"
	"github.com/nikita-reshetnyak/auth/internal/model"
	"github.com/nikita-reshetnyak/auth/internal/repository"
	"github.com/nikita-reshetnyak/auth/internal/repository/user/converter"
	modelRepo "github.com/nikita-reshetnyak/auth/internal/repository/user/model"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}
func (s *repo) Create(
	ctx context.Context,
	name string,
	email string,
	password string,
	password_confirm string,
	role model.Role,
) (int64, error) {
	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: "INSERT INTO users (name, email, password, password_confirm, role) VALUES ($1, $2, $3, $4, $5) RETURNING id",
	}
	var userId int64
	err := s.db.DB().QueryRowContext(
		ctx,
		q,
		name,
		email,
		password,
		password_confirm,
		role,
	).Scan(&userId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, fmt.Errorf("%w", err)
		}
		return 0, fmt.Errorf("pg error: %w", err)

	}
	fmt.Printf("inserted id: %d\n", userId)
	return userId, nil
}
func (s *repo) Get(ctx context.Context, id int64) (
	*model.User,
	error,
) {
	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: "SELECT id, name, email, role, created_at, updated_at FROM users WHERE id = $1",
	}
	user := modelRepo.User{}
	var roleString string
	var createdAt time.Time
	var updatedAt sql.NullTime

	err := s.db.DB().QueryRowContext(ctx,
		q, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&roleString,
		&createdAt,
		&updatedAt)
	if err != nil {
		return converter.ToUserFromRepo(&user), fmt.Errorf("%w", err)
	}
	user.Role = modelRepo.Role(roleString)
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	return converter.ToUserFromRepo(&user), nil
}
func (s *repo) Update(ctx context.Context, id int64, name string, email string) (bool, error) {
	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: "UPDATE users SET name = $1,email = $2 WHERE id = $3",
	}
	res, err := s.db.DB().ExecContext(ctx, q, name, email, id)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	fmt.Printf("updated %d rows", res.RowsAffected())
	return false, nil
}
func (s *repo) Delete(ctx context.Context, id int64) (bool, error) {
	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: "DELETE FROM users WHERE id = $1",
	}
	res, err := s.db.DB().ExecContext(ctx, q, id)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	fmt.Printf("deleted %d rows", res.RowsAffected())
	return false, nil
}
