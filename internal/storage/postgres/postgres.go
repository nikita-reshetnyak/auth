package postgres_strg

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nikita-reshetnyak/auth/internal/domains/models"
	"github.com/nikita-reshetnyak/auth/internal/storage"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Storage struct {
	conn *pgx.Conn
	ctx  context.Context
}

func New(configpath string) (*Storage, error) {
	ctx := context.Background()
	con, err := pgx.Connect(ctx, configpath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db :%w", err)
	}
	return &Storage{conn: con, ctx: ctx}, nil
}
func (s *Storage) Create(
	ctx context.Context,
	name string,
	email string,
	password string,
	password_confirm string,
	role models.UserRole,
) (int64, error) {
	var userId int64
	err := s.conn.QueryRow(
		ctx,
		"INSERT INTO users (name, email, password, password_confirm, role) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		name,
		email,
		password,
		password_confirm,
		role.String(),
	).Scan(&userId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, fmt.Errorf("%w", storage.ErrUserIsExist)
		}
		return 0, fmt.Errorf("pg error: %w", err)

	}
	fmt.Printf("inserted id: %d\n", userId)
	return userId, nil
}
func (s *Storage) Get(ctx context.Context, id int64) (
	models.User,
	error,
) {
	user := models.User{}
	var roleString string
	var createdAt time.Time
	var updatedAt time.Time
	err := s.conn.QueryRow(ctx,
		"SELECT id, name, email, role, created_at, updated_at FROM users WHERE id = $1", id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&roleString,
		&createdAt,
		&updatedAt)
	if err != nil {
		return user, fmt.Errorf("%w", err)
	}
	user.Role = models.ParseUserRole(roleString)
	user.CreatedAt = timestamppb.New(createdAt)
	user.UpdatedAt = timestamppb.New(updatedAt)
	return user, nil
}
func (s *Storage) Update(ctx context.Context, id int64, name string, email string) (*emptypb.Empty, error) {
	res, err := s.conn.Exec(ctx, "UPDATE users SET name = $1,email = $2 WHERE id = $3", name, email, id)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("%w", err)
	}
	fmt.Printf("updated %d rows", res.RowsAffected())
	return &emptypb.Empty{}, nil
}
func (s *Storage) Delete(ctx context.Context, id int64) (*emptypb.Empty, error) {
	res, err := s.conn.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("%w", err)
	}
	fmt.Printf("deleted %d rows", res.RowsAffected())
	return &emptypb.Empty{}, nil
}
