package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/qreasio/go-starter-kit/internal/user"
	"github.com/qreasio/go-starter-kit/pkg/log"
	"github.com/qreasio/go-starter-kit/pkg/model"
)

func TestUserRepository_ListUsers(t *testing.T) {
	logger := log.New()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := sqlx.NewDb(mockDB, "sqlmock")
	defer mockDB.Close()

	repo := NewRepository(db, logger)

	t.Log("Test List Users with default parameter")
	{
		rows := sqlmock.NewRows([]string{"first_name", "last_name", "email", "date_joined", "last_login"}).
			AddRow("Isak", "Rickyanto", "isak.rickyanto@gmail.com", time.Now(), time.Now()).
			AddRow("Farida", "Tjandra", "farida@tjandra.com", time.Now(), time.Now())
		mock.ExpectQuery(ListUsersSQL).WithArgs(0, 2).WillReturnRows(rows)

		p := model.NewPagination()
		users, err := repo.List(context.Background(), &ListUsersRequest{Pagination: *p, Search: ""})

		if err != nil {
			t.Errorf("Failed to get list of users %s", err)
		}

		if len(users) != p.Limit {
			t.Errorf("Failed to get default number of list users correctly")
		}

	}

	t.Log("Test List Users with specific limit")
	{
		row := sqlmock.NewRows([]string{"first_name", "last_name", "email", "date_joined", "last_login"}).
			AddRow("Isak", "Rickyanto", "isak.rickyanto@gmail.com", time.Now(), time.Now())
		mock.ExpectQuery(ListUsersSQL).WithArgs(0, 1).WillReturnRows(row)

		p := model.NewPagination()
		p.Limit = 1
		users, err := repo.List(context.Background(), &ListUsersRequest{Pagination: *p, Search: ""})

		if err != nil {
			t.Errorf("Failed to get list of users %s", err)
		}

		if len(users) != 1 {
			t.Errorf("Failed to get number of list users correctly")
		}

	}
}
