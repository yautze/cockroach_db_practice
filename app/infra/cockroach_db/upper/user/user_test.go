package user

import (
	"cockroach-practice/app/infra/config"
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// TestList -
func TestList(t *testing.T) {
	repo, err := NewRepo(config.C)
	if err != nil {
		t.Fatal(err)
	}

	users, err := repo.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	spew.Dump(users)
}
