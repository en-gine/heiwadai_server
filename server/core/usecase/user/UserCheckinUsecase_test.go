package user

import (
	"context"
	"testing"
	"time"

	"server/core/entity"
	"server/core/errors"
	queryService "server/core/infra/queryService"
	"server/core/infra/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserQueryService
type MockUserQueryService struct {
	mock.Mock
}

func (m *MockUserQueryService) GetByID(authID uuid.UUID) (*entity.User, error) {
	args := m.Called(authID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserQueryService) GetAll() ([]*entity.User, error) {
	args := m.Called()
	return args.Get(0).([]*entity.User), args.Error(1)
}

func (m *MockUserQueryService) GetByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserQueryService) GetByIDForUpdate(authID uuid.UUID, tx interface{}) (*entity.User, error) {
	args := m.Called(authID, tx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

// MockStoreRepository
type MockStoreRepository struct {
	mock.Mock
}

func (m *MockStoreRepository) Save(tx repository.ITransaction, store *entity.Store) error {
	args := m.Called(tx, store)
	return args.Error(0)
}

func (m *MockStoreRepository) Update(tx repository.ITransaction, store *entity.Store) error {
	args := m.Called(tx, store)
	return args.Error(0)
}

func (m *MockStoreRepository) UpdateUnlimitQrCode(tx repository.ITransaction, storeID uuid.UUID, unlimitQrCodeHash uuid.UUID) error {
	args := m.Called(tx, storeID, unlimitQrCodeHash)
	return args.Error(0)
}

func (m *MockStoreRepository) UpdateQrCode(tx repository.ITransaction, storeID uuid.UUID, qrCodeHash uuid.UUID) error {
	args := m.Called(tx, storeID, qrCodeHash)
	return args.Error(0)
}

func (m *MockStoreRepository) Delete(tx repository.ITransaction, storeID uuid.UUID) error {
	args := m.Called(tx, storeID)
	return args.Error(0)
}

// MockCheckinRepository
type MockCheckinRepository struct {
	mock.Mock
}

func (m *MockCheckinRepository) Save(tx repository.ITransaction, checkin *entity.Checkin) error {
	args := m.Called(tx, checkin)
	return args.Error(0)
}

func (m *MockCheckinRepository) BulkArchive(tx repository.ITransaction, userID uuid.UUID) error {
	args := m.Called(tx, userID)
	return args.Error(0)
}

// MockCouponRepository
type MockCouponRepository struct {
	mock.Mock
}

func (m *MockCouponRepository) Save(tx repository.ITransaction, coupon interface{}) error {
	args := m.Called(tx, coupon)
	return args.Error(0)
}

func (m *MockCouponRepository) Update(tx repository.ITransaction, coupon interface{}) error {
	args := m.Called(tx, coupon)
	return args.Error(0)
}

func (m *MockCouponRepository) Delete(tx repository.ITransaction, couponID uuid.UUID) error {
	args := m.Called(tx, couponID)
	return args.Error(0)
}

// MockUserCouponRepository
type MockUserCouponRepository struct {
	mock.Mock
}

func (m *MockUserCouponRepository) Save(tx repository.ITransaction, userCoupon *entity.UserAttachedCoupon) error {
	args := m.Called(tx, userCoupon)
	return args.Error(0)
}

func (m *MockUserCouponRepository) Update(tx repository.ITransaction, userCoupon *entity.UserAttachedCoupon) error {
	args := m.Called(tx, userCoupon)
	return args.Error(0)
}

// MockUserCouponQueryService
type MockUserCouponQueryService struct {
	mock.Mock
}

func (m *MockUserCouponQueryService) GetMyNotUsedByID(authID uuid.UUID, couponID uuid.UUID) (*entity.UserAttachedCoupon, error) {
	args := m.Called(authID, couponID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UserAttachedCoupon), args.Error(1)
}

func (m *MockUserCouponQueryService) GetMyAvailable(authID uuid.UUID) ([]*entity.UserAttachedCoupon, error) {
	args := m.Called(authID)
	return args.Get(0).([]*entity.UserAttachedCoupon), args.Error(1)
}

func (m *MockUserCouponQueryService) GetMyUsed(authID uuid.UUID) ([]*entity.UserAttachedCoupon, error) {
	args := m.Called(authID)
	return args.Get(0).([]*entity.UserAttachedCoupon), args.Error(1)
}

func (m *MockUserCouponQueryService) GetMyExpired(authID uuid.UUID) ([]*entity.UserAttachedCoupon, error) {
	args := m.Called(authID)
	return args.Get(0).([]*entity.UserAttachedCoupon), args.Error(1)
}

// MockStoreQueryService
type MockStoreQueryService struct {
	mock.Mock
}

func (m *MockStoreQueryService) GetByID(storeID uuid.UUID) (*entity.Store, error) {
	args := m.Called(storeID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Store), args.Error(1)
}

func (m *MockStoreQueryService) GetAll() ([]*entity.Store, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Store), args.Error(1)
}

func (m *MockStoreQueryService) GetActiveAll() ([]*entity.Store, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Store), args.Error(1)
}

func (m *MockStoreQueryService) GetStoreByQrCode(qrCode uuid.UUID) (*entity.Store, error) {
	args := m.Called(qrCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Store), args.Error(1)
}

func (m *MockStoreQueryService) GetStoreByUnlimitQrCode(qrCode uuid.UUID) (*entity.Store, error) {
	args := m.Called(qrCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Store), args.Error(1)
}

// MockCheckinQueryService
type MockCheckinQueryService struct {
	mock.Mock
}

func (m *MockCheckinQueryService) GetMyActiveCheckin(authID uuid.UUID) ([]*entity.Checkin, error) {
	args := m.Called(authID)
	return args.Get(0).([]*entity.Checkin), args.Error(1)
}

func (m *MockCheckinQueryService) GetMyLastStoreCheckin(authID uuid.UUID, storeID uuid.UUID) (*entity.Checkin, error) {
	args := m.Called(authID, storeID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Checkin), args.Error(1)
}

// MockCouponQueryService
type MockCouponQueryService struct {
	mock.Mock
}

func (m *MockCouponQueryService) GetByIDForUpdate(couponID uuid.UUID, tx interface{}) (*entity.Coupon, error) {
	args := m.Called(couponID, tx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Coupon), args.Error(1)
}

func (m *MockCouponQueryService) GetAll() ([]*entity.Coupon, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Coupon), args.Error(1)
}

func (m *MockCouponQueryService) GetAllActive() ([]*entity.Coupon, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Coupon), args.Error(1)
}

func (m *MockCouponQueryService) GetMyAll(authID uuid.UUID) ([]*entity.Coupon, error) {
	args := m.Called(authID)
	return args.Get(0).([]*entity.Coupon), args.Error(1)
}

func (m *MockCouponQueryService) GetAvailableByStoreID(storeID uuid.UUID) ([]*entity.Coupon, error) {
	args := m.Called(storeID)
	return args.Get(0).([]*entity.Coupon), args.Error(1)
}

func (m *MockCouponQueryService) GetByID(couponID uuid.UUID) (*entity.Coupon, error) {
	args := m.Called(couponID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Coupon), args.Error(1)
}

// MockTransaction
type MockTransaction struct {
	mock.Mock
}

func (m *MockTransaction) Begin(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockTransaction) Commit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockTransaction) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockTransaction) Tran() interface{} {
	args := m.Called()
	return args.Get(0)
}

// MockMemoryRepository
type MockMemoryRepository struct {
	mock.Mock
}

func (m *MockMemoryRepository) Get(key string) *[]byte {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*[]byte)
}

func (m *MockMemoryRepository) Set(key string, value []byte, expire time.Duration) {
	m.Called(key, value, expire)
}

func (m *MockMemoryRepository) Delete(key string) {
	m.Called(key)
}

func (m *MockMemoryRepository) SetNX(key string, value []byte, expire time.Duration) bool {
	args := m.Called(key, value, expire)
	return args.Bool(0)
}

func TestCheckin_Success(t *testing.T) {
	// Setup
	mockUserQuery := new(MockUserQueryService)
	mockStoreRepo := new(MockStoreRepository)
	mockCheckinRepo := new(MockCheckinRepository)
	mockCouponRepo := new(MockCouponRepository)
	mockUserCouponRepo := new(MockUserCouponRepository)
	mockUserCouponQuery := new(MockUserCouponQueryService)
	mockStoreQuery := new(MockStoreQueryService)
	mockCheckinQuery := new(MockCheckinQueryService)
	mockCouponQuery := new(MockCouponQueryService)
	mockTransaction := new(MockTransaction)
	mockMemoryRepo := new(MockMemoryRepository)

	usecase := NewUserCheckinUsecase(
		mockUserQuery,
		mockStoreRepo,
		mockCheckinRepo,
		mockCouponRepo,
		mockUserCouponRepo,
		mockUserCouponQuery,
		mockStoreQuery,
		mockCheckinQuery,
		mockCouponQuery,
		mockTransaction,
		mockMemoryRepo,
	)

	authID := uuid.New()
	qrHash := uuid.New()
	storeID := uuid.New()

	user := &entity.User{
		ID:        authID,
		FirstName: "Test",
		LastName:  "User",
	}

	store := &entity.Store{
		ID:       storeID,
		Name:     "Test Store",
		IsActive: true,
	}

	// Mock expectations
	mockMemoryRepo.On("SetNX", mock.AnythingOfType("string"), []byte("processing"), 5*time.Second).Return(true)
	mockMemoryRepo.On("Delete", mock.AnythingOfType("string")).Return()

	mockUserQuery.On("GetByID", authID).Return(user, nil)
	mockStoreQuery.On("GetStoreByQrCode", qrHash).Return(store, nil)
	mockStoreQuery.On("GetStoreByUnlimitQrCode", qrHash).Return(nil, nil)
	mockCheckinQuery.On("GetMyLastStoreCheckin", authID, storeID).Return(nil, nil)
	mockTransaction.On("Begin", mock.Anything).Return(nil)
	mockTransaction.On("Tran").Return(&struct{}{})
	mockCheckinQuery.On("GetMyActiveCheckin", authID).Return([]*entity.Checkin{}, nil)
	mockCheckinRepo.On("Save", mockTransaction, mock.AnythingOfType("*entity.Checkin")).Return(nil)
	mockTransaction.On("Commit").Return(nil)
	mockCheckinQuery.On("GetMyActiveCheckin", authID).Return([]*entity.Checkin{
		{
			ID:        uuid.New(),
			Store:     store,
			User:      user,
			CheckInAt: time.Now(),
			Archive:   false,
		},
	}, nil)

	// Execute
	stampCard, coupon, err := usecase.Checkin(authID, qrHash)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, stampCard)
	assert.Nil(t, coupon)
	assert.Equal(t, 1, stampCard.CheckinCount)

	// Verify mock expectations
	mockMemoryRepo.AssertExpectations(t)
	mockUserQuery.AssertExpectations(t)
	mockStoreQuery.AssertExpectations(t)
	mockCheckinQuery.AssertExpectations(t)
	mockCheckinRepo.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
}

func TestCheckin_DuplicateRequest(t *testing.T) {
	// Setup
	mockUserQuery := new(MockUserQueryService)
	mockStoreRepo := new(MockStoreRepository)
	mockCheckinRepo := new(MockCheckinRepository)
	mockCouponRepo := new(MockCouponRepository)
	mockUserCouponRepo := new(MockUserCouponRepository)
	mockUserCouponQuery := new(MockUserCouponQueryService)
	mockStoreQuery := new(MockStoreQueryService)
	mockCheckinQuery := new(MockCheckinQueryService)
	mockCouponQuery := new(MockCouponQueryService)
	mockTransaction := new(MockTransaction)
	mockMemoryRepo := new(MockMemoryRepository)

	usecase := NewUserCheckinUsecase(
		mockUserQuery,
		mockStoreRepo,
		mockCheckinRepo,
		mockCouponRepo,
		mockUserCouponRepo,
		mockUserCouponQuery,
		mockStoreQuery,
		mockCheckinQuery,
		mockCouponQuery,
		mockTransaction,
		mockMemoryRepo,
	)

	authID := uuid.New()
	qrHash := uuid.New()

	// Mock expectations - SetNX returns false indicating key already exists
	mockMemoryRepo.On("SetNX", mock.AnythingOfType("string"), []byte("processing"), 5*time.Second).Return(false)

	// Execute
	stampCard, coupon, err := usecase.Checkin(authID, qrHash)

	// Assert
	assert.Nil(t, stampCard)
	assert.Nil(t, coupon)
	assert.NotNil(t, err)
	assert.Equal(t, errors.UnPemitedOperation, err.Type)
	assert.Contains(t, err.Error(), "チェックイン処理中です")

	// Verify mock expectations
	mockMemoryRepo.AssertExpectations(t)
}

func TestCheckin_Within24Hours(t *testing.T) {
	// Setup
	mockUserQuery := new(MockUserQueryService)
	mockStoreRepo := new(MockStoreRepository)
	mockCheckinRepo := new(MockCheckinRepository)
	mockCouponRepo := new(MockCouponRepository)
	mockUserCouponRepo := new(MockUserCouponRepository)
	mockUserCouponQuery := new(MockUserCouponQueryService)
	mockStoreQuery := new(MockStoreQueryService)
	mockCheckinQuery := new(MockCheckinQueryService)
	mockCouponQuery := new(MockCouponQueryService)
	mockTransaction := new(MockTransaction)
	mockMemoryRepo := new(MockMemoryRepository)

	usecase := NewUserCheckinUsecase(
		mockUserQuery,
		mockStoreRepo,
		mockCheckinRepo,
		mockCouponRepo,
		mockUserCouponRepo,
		mockUserCouponQuery,
		mockStoreQuery,
		mockCheckinQuery,
		mockCouponQuery,
		mockTransaction,
		mockMemoryRepo,
	)

	authID := uuid.New()
	qrHash := uuid.New()
	storeID := uuid.New()

	user := &entity.User{
		ID:        authID,
		FirstName: "Test",
		LastName:  "User",
	}

	store := &entity.Store{
		ID:       storeID,
		Name:     "Test Store",
		IsActive: true,
	}

	lastCheckin := &entity.Checkin{
		ID:        uuid.New(),
		Store:     store,
		User:      user,
		CheckInAt: time.Now().Add(-1 * time.Hour), // 1 hour ago
		Archive:   false,
	}

	// Mock expectations
	mockMemoryRepo.On("SetNX", mock.AnythingOfType("string"), []byte("processing"), 5*time.Second).Return(true)
	mockMemoryRepo.On("Delete", mock.AnythingOfType("string")).Return()

	mockUserQuery.On("GetByID", authID).Return(user, nil)
	mockStoreQuery.On("GetStoreByQrCode", qrHash).Return(store, nil)
	mockStoreQuery.On("GetStoreByUnlimitQrCode", qrHash).Return(nil, nil)
	mockCheckinQuery.On("GetMyLastStoreCheckin", authID, storeID).Return(lastCheckin, nil)

	// Execute
	stampCard, coupon, err := usecase.Checkin(authID, qrHash)

	// Assert
	assert.Nil(t, stampCard)
	assert.Nil(t, coupon)
	assert.NotNil(t, err)
	assert.Equal(t, errors.UnPemitedOperation, err.Type)
	assert.Contains(t, err.Error(), "24時間以内にチェックインした店舗はチェックインできません")

	// Verify mock expectations
	mockMemoryRepo.AssertExpectations(t)
	mockUserQuery.AssertExpectations(t)
	mockStoreQuery.AssertExpectations(t)
	mockCheckinQuery.AssertExpectations(t)
}

func TestCheckin_CouponGeneration(t *testing.T) {
	// Setup
	mockUserQuery := new(MockUserQueryService)
	mockStoreRepo := new(MockStoreRepository)
	mockCheckinRepo := new(MockCheckinRepository)
	mockCouponRepo := new(MockCouponRepository)
	mockUserCouponRepo := new(MockUserCouponRepository)
	mockUserCouponQuery := new(MockUserCouponQueryService)
	mockStoreQuery := new(MockStoreQueryService)
	mockCheckinQuery := new(MockCheckinQueryService)
	mockCouponQuery := new(MockCouponQueryService)
	mockTransaction := new(MockTransaction)
	mockMemoryRepo := new(MockMemoryRepository)

	usecase := NewUserCheckinUsecase(
		mockUserQuery,
		mockStoreRepo,
		mockCheckinRepo,
		mockCouponRepo,
		mockUserCouponRepo,
		mockUserCouponQuery,
		mockStoreQuery,
		mockCheckinQuery,
		mockCouponQuery,
		mockTransaction,
		mockMemoryRepo,
	)

	authID := uuid.New()
	qrHash := uuid.New()
	storeID := uuid.New()

	user := &entity.User{
		ID:        authID,
		FirstName: "Test",
		LastName:  "User",
	}

	store := &entity.Store{
		ID:       storeID,
		Name:     "Test Store",
		IsActive: true,
	}

	// Create 4 existing checkins
	existingCheckins := make([]*entity.Checkin, 4)
	for i := 0; i < 4; i++ {
		existingCheckins[i] = &entity.Checkin{
			ID:        uuid.New(),
			Store:     store,
			User:      user,
			CheckInAt: time.Now(),
			Archive:   false,
		}
	}

	allStores := []*entity.Store{store}

	// Mock expectations
	mockMemoryRepo.On("SetNX", mock.AnythingOfType("string"), []byte("processing"), 5*time.Second).Return(true)
	mockMemoryRepo.On("Delete", mock.AnythingOfType("string")).Return()

	mockUserQuery.On("GetByID", authID).Return(user, nil)
	mockStoreQuery.On("GetStoreByQrCode", qrHash).Return(store, nil)
	mockStoreQuery.On("GetStoreByUnlimitQrCode", qrHash).Return(nil, nil)
	mockCheckinQuery.On("GetMyLastStoreCheckin", authID, storeID).Return(nil, nil)
	mockTransaction.On("Begin", mock.Anything).Return(nil)
	mockTransaction.On("Tran").Return(&struct{}{})
	mockCheckinQuery.On("GetMyActiveCheckin", authID).Return(existingCheckins, nil)
	mockCheckinRepo.On("Save", mockTransaction, mock.AnythingOfType("*entity.Checkin")).Return(nil)
	mockStoreQuery.On("GetActiveAll").Return(allStores, nil)
	mockCouponRepo.On("Save", mockTransaction, mock.AnythingOfType("*entity.StandardCoupon")).Return(nil)
	mockUserCouponRepo.On("Save", mockTransaction, mock.AnythingOfType("*entity.UserAttachedCoupon")).Return(nil)
	mockCouponRepo.On("Save", mockTransaction, mock.AnythingOfType("*entity.IssuedCoupon")).Return(nil)
	mockCheckinRepo.On("BulkArchive", mockTransaction, authID).Return(nil)
	mockTransaction.On("Commit").Return(nil)
	mockCheckinQuery.On("GetMyActiveCheckin", authID).Return([]*entity.Checkin{}, nil) // After archive

	// Execute
	stampCard, coupon, err := usecase.Checkin(authID, qrHash)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, stampCard)
	assert.NotNil(t, coupon)
	assert.Equal(t, 0, stampCard.CheckinCount) // Reset after coupon generation

	// Verify mock expectations
	mockMemoryRepo.AssertExpectations(t)
	mockUserQuery.AssertExpectations(t)
	mockStoreQuery.AssertExpectations(t)
	mockCheckinQuery.AssertExpectations(t)
	mockCheckinRepo.AssertExpectations(t)
	mockCouponRepo.AssertExpectations(t)
	mockUserCouponRepo.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
}