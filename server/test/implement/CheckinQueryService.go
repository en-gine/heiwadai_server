package test

import (
	"testing"

	"server/core/entity"
	"server/core/infra/queryService/types"
	"server/infrastructure/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

// Mocked user and store entity for testing
var (
	user  = &entity.User{ID: uuid.New()}
	store = &entity.Store{ID: uuid.New()}
)

var limit = 10

var pager = types.NewPageQuery(nil, &limit)

func TestCheckinQueryService(t *testing.T) {
	_, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	checkinQueryService := &repository.CheckinQueryService{}

	mock.ExpectQuery("^SELECT (.+) FROM checkins").WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "store_id"}))

	t.Run("GetActiveCheckin", func(t *testing.T) {
		_, err := checkinQueryService.GetActiveCheckin(user.ID)
		if err != nil {
			t.Error("Error occurred while trying to get active checkins")
		}
	})

	t.Run("GetLastStoreCheckin", func(t *testing.T) {
		_, err := checkinQueryService.GetLastStoreCheckin(user.ID, store.ID)
		if err != nil {
			t.Error("Error occurred while trying to get last store checkin")
		}
	})

	t.Run("GetAllCheckin", func(t *testing.T) {
		_, err := checkinQueryService.GetAllCheckin(user.ID, pager)
		if err != nil {
			t.Error("Error occurred while trying to get all checkins")
		}
	})
}
