package user_test

import (
	"context"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/qreasio/go-starter-kit/internal/user"
	"github.com/qreasio/go-starter-kit/pkg/log"
	"github.com/qreasio/go-starter-kit/pkg/test"
)

var (
	db     *sqlx.DB
	logger = log.New()
)

func TestMain(m *testing.M) {
	var err error
	var terminateContainer func() // variable to store function to terminate container
	terminateContainer, db, err = test.SetupMySQLContainer(logger)
	defer terminateContainer() // make sure container will be terminated at the end
	if err != nil {
		logger.Error("failed to setup MySQL container")
		panic(err)
	}
	os.Exit(m.Run())
}

func TestUserRepository_ListIntegration(t *testing.T) {
	repo := user.NewRepository(db, logger)
	ctx := context.Background()
	req := user.NewListUsersRequest()
	users, err := repo.List(ctx, &req)

	if err != nil {
		t.Errorf("error on list users : %s", err)
	}

	if len(users) < 1 {
		t.Errorf("Failed to get list users : %s", err)
	}

	want := "isak"
	got := users[0].Username
	if got != want {
		t.Errorf("Error get user, want : %s, got : %s", want, got)
	}
}
