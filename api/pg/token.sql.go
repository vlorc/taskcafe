// Code generated by sqlc. DO NOT EDIT.
// source: token.sql

package pg

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createRefreshToken = `-- name: CreateRefreshToken :one
INSERT INTO refresh_token (user_id, created_at, expires_at) VALUES ($1, $2, $3) RETURNING token_id, user_id, created_at, expires_at
`

type CreateRefreshTokenParams struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (q *Queries) CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, createRefreshToken, arg.UserID, arg.CreatedAt, arg.ExpiresAt)
	var i RefreshToken
	err := row.Scan(
		&i.TokenID,
		&i.UserID,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const deleteExpiredTokens = `-- name: DeleteExpiredTokens :exec
DELETE FROM refresh_token WHERE expires_at <= NOW()
`

func (q *Queries) DeleteExpiredTokens(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteExpiredTokens)
	return err
}

const deleteRefreshTokenByID = `-- name: DeleteRefreshTokenByID :exec
DELETE FROM refresh_token WHERE token_id = $1
`

func (q *Queries) DeleteRefreshTokenByID(ctx context.Context, tokenID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteRefreshTokenByID, tokenID)
	return err
}

const deleteRefreshTokenByUserID = `-- name: DeleteRefreshTokenByUserID :exec
DELETE FROM refresh_token WHERE user_id = $1
`

func (q *Queries) DeleteRefreshTokenByUserID(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteRefreshTokenByUserID, userID)
	return err
}

const getRefreshTokenByID = `-- name: GetRefreshTokenByID :one
SELECT token_id, user_id, created_at, expires_at FROM refresh_token WHERE token_id = $1
`

func (q *Queries) GetRefreshTokenByID(ctx context.Context, tokenID uuid.UUID) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, getRefreshTokenByID, tokenID)
	var i RefreshToken
	err := row.Scan(
		&i.TokenID,
		&i.UserID,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}