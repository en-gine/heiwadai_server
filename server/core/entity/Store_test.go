package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateStore(t *testing.T) {
	t.Run("通常の店舗作成", func(t *testing.T) {
		id := uuid.New()
		name := "Test Store"
		branchName := "Branch 1"
		store, err := CreateStore(id, name, &branchName, "123-4567", "Tokyo", "03-1234-5678", "http://example.com", "http://example.com/stamp.png", false, nil)

		assert.Nil(t, err)
		assert.NotNil(t, store)
		assert.Equal(t, id, store.ID)
		assert.Equal(t, name, store.Name)
		assert.Equal(t, &branchName, store.BranchName)
		assert.False(t, store.Stayable)
		assert.True(t, store.IsActive)
		assert.NotEqual(t, uuid.Nil, store.QRCode)
		assert.NotEqual(t, uuid.Nil, store.UnLimitedQRCode)
	})

	t.Run("宿泊可能な店舗作成", func(t *testing.T) {
		id := uuid.New()
		stayableInfo := CreateStayableStoreInfo("Available", 35.6895, 139.6917, "Train", "http://api.example.com", "BOOKING123")
		store, err := CreateStore(id, "Stayable Store", nil, "123-4567", "Tokyo", "03-1234-5678", "http://example.com", "http://example.com/stamp.png", true, stayableInfo)

		assert.Nil(t, err)
		assert.NotNil(t, store)
		assert.True(t, store.Stayable)
		assert.Equal(t, stayableInfo, store.StayableStoreInfo)
	})

	t.Run("宿泊可能だが情報がない場合", func(t *testing.T) {
		id := uuid.New()
		store, err := CreateStore(id, "Invalid Store", nil, "123-4567", "Tokyo", "03-1234-5678", "http://example.com", "http://example.com/stamp.png", true, nil)

		assert.Nil(t, store)
		assert.NotNil(t, err)
		assert.Equal(t, "宿泊可能な施設は宿泊施設情報を入力してください。", err.Error())
	})
}

func TestCreateStayableStoreInfo(t *testing.T) {
	info := CreateStayableStoreInfo("Available", 35.6895, 139.6917, "Train", "http://api.example.com", "BOOKING123")

	assert.NotNil(t, info)
	assert.Equal(t, "Available", info.Parking)
	assert.Equal(t, 35.6895, info.Latitude)
	assert.Equal(t, 139.6917, info.Longitude)
	assert.Equal(t, "Train", info.AccessInfo)
	assert.Equal(t, "http://api.example.com", info.RestAPIURL)
	assert.Equal(t, "BOOKING123", info.BookingSystemID)
}

func TestCreateStayableStore(t *testing.T) {
	t.Run("正常な宿泊可能店舗作成", func(t *testing.T) {
		store, _ := CreateStore(uuid.New(), "Stayable Store", nil, "123-4567", "Tokyo", "03-1234-5678", "http://example.com", "http://example.com/stamp.png", true, CreateStayableStoreInfo("Available", 35.6895, 139.6917, "Train", "http://api.example.com", "BOOKING123"))
		stayableStore, err := CreateStayableStore(store, store.StayableStoreInfo)

		assert.Nil(t, err)
		assert.NotNil(t, stayableStore)
		assert.Equal(t, store, stayableStore.Store)
		assert.Equal(t, store.StayableStoreInfo, stayableStore.StayableStoreInfo)
	})

	t.Run("Storeがnilの場合", func(t *testing.T) {
		stayableStore, err := CreateStayableStore(nil, CreateStayableStoreInfo("Available", 35.6895, 139.6917, "Train", "http://api.example.com", "BOOKING123"))

		assert.Nil(t, stayableStore)
		assert.NotNil(t, err)
		assert.Equal(t, "Storeがnilです。", err.Error())
	})

	t.Run("StayableInfoがnilの場合", func(t *testing.T) {
		store, _ := CreateStore(uuid.New(), "Store", nil, "123-4567", "Tokyo", "03-1234-5678", "http://example.com", "http://example.com/stamp.png", true, CreateStayableStoreInfo("Available", 35.6895, 139.6917, "Train", "http://api.example.com", "BOOKING123"))
		stayableStore, err := CreateStayableStore(store, nil)

		assert.Nil(t, stayableStore)
		assert.NotNil(t, err)
		assert.Equal(t, "StayableInfoがnilです。", err.Error())
	})

	t.Run("Stayableがfalseの場合", func(t *testing.T) {
		store, _ := CreateStore(uuid.New(), "Non-Stayable Store", nil, "123-4567", "Tokyo", "03-1234-5678", "http://example.com", "http://example.com/stamp.png", false, nil)
		stayableStore, err := CreateStayableStore(store, CreateStayableStoreInfo("Available", 35.6895, 139.6917, "Train", "http://api.example.com", "BOOKING123"))

		assert.Nil(t, stayableStore)
		assert.NotNil(t, err)
		assert.Equal(t, "Stayableがtrueである必要があります。", err.Error())
	})
}

func TestRegenStore(t *testing.T) {
	id := uuid.New()
	qrCode := uuid.New()
	unlimitedQRCode := uuid.New()
	branchName := "Branch 1"
	stayableInfo := CreateStayableStoreInfo("Available", 35.6895, 139.6917, "Train", "http://api.example.com", "BOOKING123")

	store := RegenStore(id, "Regen Store", &branchName, "123-4567", "Tokyo", "03-1234-5678", "http://example.com", "http://example.com/stamp.png", true, true, qrCode, unlimitedQRCode, stayableInfo)

	assert.NotNil(t, store)
	assert.Equal(t, id, store.ID)
	assert.Equal(t, "Regen Store", store.Name)
	assert.Equal(t, &branchName, store.BranchName)
	assert.Equal(t, "123-4567", store.ZipCode)
	assert.Equal(t, "Tokyo", store.Address)
	assert.Equal(t, "03-1234-5678", store.Tel)
	assert.Equal(t, "http://example.com", store.SiteURL)
	assert.Equal(t, "http://example.com/stamp.png", store.StampImageURL)
	assert.True(t, store.Stayable)
	assert.True(t, store.IsActive)
	assert.Equal(t, qrCode, store.QRCode)
	assert.Equal(t, unlimitedQRCode, store.UnLimitedQRCode)
	assert.Equal(t, stayableInfo, store.StayableStoreInfo)
}

func TestRegenStayableStoreInfo(t *testing.T) {
	info := RegenStayableStoreInfo("Available", 35.6895, 139.6917, "Train", "http://api.example.com", "BOOKING123")

	assert.NotNil(t, info)
	assert.Equal(t, "Available", info.Parking)
	assert.Equal(t, 35.6895, info.Latitude)
	assert.Equal(t, 139.6917, info.Longitude)
	assert.Equal(t, "Train", info.AccessInfo)
	assert.Equal(t, "http://api.example.com", info.RestAPIURL)
	assert.Equal(t, "BOOKING123", info.BookingSystemID)
}

func TestRegenStayableStore(t *testing.T) {
	store := RegenStore(uuid.New(), "Regen Stayable Store", nil, "123-4567", "Tokyo", "03-1234-5678", "http://example.com", "http://example.com/stamp.png", true, true, uuid.New(), uuid.New(), nil)
	stayableInfo := RegenStayableStoreInfo("Available", 35.6895, 139.6917, "Train", "http://api.example.com", "BOOKING123")

	stayableStore := RegenStayableStore(store, stayableInfo)

	assert.NotNil(t, stayableStore)
	assert.Equal(t, store, stayableStore.Store)
	assert.Equal(t, stayableInfo, stayableStore.StayableStoreInfo)
}
