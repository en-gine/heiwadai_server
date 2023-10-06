package queryservice

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IStoreQueryService interface {
	GetByID(id uuid.UUID) (*entity.Store, error)
	GetActiveAll() ([]*entity.Store, error)
	GetStayables() ([]*entity.StayableStore, error)
	GetStayableByID(id uuid.UUID) (*entity.StayableStore, error)
	GetAll() ([]*entity.Store, error)                                    // activeかどうか不問
	GetStayableByBookingID(bookID string) (*entity.StayableStore, error) // bookingIDからstayableなstoreを取得
}
