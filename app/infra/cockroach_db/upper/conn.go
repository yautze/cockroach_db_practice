package upper

import (
	"cockroach-practice/app/infra/config"
	"time"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
)

// NewWithUpper -
func NewWithUpper(c config.Config) (db.Session, error) {
	settings := cockroachdb.ConnectionURL{
		Host:     c.Cockroach.Host,
		User:     c.Cockroach.User,
		Password: c.Cockroach.Password,
		Database: c.Cockroach.Database,
		Options: map[string]string{
			"sslmode": "disable",
		},
	}

	// open db connection
	session, err := cockroachdb.Open(settings)
	if err != nil {
		return nil, err
	}

	// set connection pool
	session.SetMaxOpenConns(10)
	session.SetConnMaxLifetime(60 * time.Second)
	session.SetMaxIdleConns(3)

	return session, nil
}
