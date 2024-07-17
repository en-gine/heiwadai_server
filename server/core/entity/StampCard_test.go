package entity

import (
	"testing"
	"time"

	"server/core/errors"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewStampCard(t *testing.T) {
	t.Run("正常系: 有効なチェックイン数", func(t *testing.T) {
		checkIns := createMockCheckIns(3)
		stampCard, err := NewStampCard(checkIns)

		assert.Nil(t, err)
		assert.NotNil(t, stampCard)
		assert.Len(t, stampCard.Stamps, 3)

		for i, stamp := range stampCard.Stamps {
			assert.Equal(t, checkIns[i].ID, *stamp.CheckinID)
			assert.Equal(t, checkIns[i].Store.Name, *stamp.StoreName)
			assert.Equal(t, checkIns[i].Store.ID, *stamp.StoreID)
			assert.Equal(t, checkIns[i].Store.StampImageURL, *stamp.StoreStampImage)
			assert.Equal(t, checkIns[i].CheckInAt, *stamp.CheckInAt)
		}
	})

	t.Run("異常系: チェックイン数が上限を超える", func(t *testing.T) {
		checkIns := createMockCheckIns(MaxStampCount + 1)
		stampCard, err := NewStampCard(checkIns)

		assert.Nil(t, stampCard)
		assert.NotNil(t, err)
		assert.IsType(t, &errors.DomainError{}, err)
		assert.Equal(t, "チェックイン数が上限を超えています。", err.Error())
	})

	t.Run("正常系: チェックインがない", func(t *testing.T) {
		checkIns := []*Checkin{}
		stampCard, err := NewStampCard(checkIns)

		assert.Nil(t, err)
		assert.NotNil(t, stampCard)
		assert.Len(t, stampCard.Stamps, 0)
	})
}

// テスト用のモックCheckinデータを生成する補助関数
func createMockCheckIns(count int) []*Checkin {
	checkIns := make([]*Checkin, count)
	for i := 0; i < count; i++ {
		checkIns[i] = &Checkin{
			ID: uuid.New(),
			Store: &Store{
				ID:            uuid.New(),
				Name:          "Store " + string(rune(i+'A')),
				StampImageURL: "https://example.com/stamp" + string(rune(i+'A')) + ".png",
			},
			CheckInAt: time.Now().Add(time.Duration(i) * time.Hour),
		}
	}
	return checkIns
}
