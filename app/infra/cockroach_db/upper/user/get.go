package user

import (
	"cockroach-practice/app/domain/model"
	"context"
	"log"
)

// List -
func (r *repo) List(ctx context.Context) ([]*model.User, error) {
	sql := r.db.SQL().Select("*").From(C)

	users := make([]*model.User, 0)
	if err := sql.All(&users); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return users, nil
}
