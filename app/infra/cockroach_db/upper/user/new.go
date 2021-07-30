package user

import (
	"cockroach-practice/app/domain/repository"
	u "cockroach-practice/app/infra/cockroach_db/upper"
	"cockroach-practice/app/infra/config"

	"github.com/upper/db/v4"
)

// C -
const C = "user"

// repo -
type repo struct {
	db           db.Session
	databaseName string
}

// NewRepo -
func NewRepo(c config.Config) (repository.UserRepository, error) {
	d, err := u.NewWithUpper(c)
	if err != nil {
		return nil, err
	}
	return &repo{
		db:           d,
		databaseName: c.Cockroach.Database,
	}, err
}
