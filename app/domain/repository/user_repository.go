package repository

import (
	"cockroach-practice/app/domain/model"
	"context"
)

// UserRepository -
type UserRepository interface {
	List(ctx context.Context) ([]*model.User, error)
}
