package queryservice

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IStoreQueryService interface {
	GetByID(id uuid.UUID) (*entity.Store, error)
	GetActiveAll() ([]*entity.Store, error)
	GetStayables() ([]*entity.StayableStore, error)
	GetStayableByID(id uuid.UUID) (*entity.StayableStore, error)         // stayableなstoreを取得実際にstayableinfoが入ってるかどうかは不問
	GetAll() ([]*entity.Store, error)                                    // activeかどうか不問
	GetStayableByBookingID(bookID string) (*entity.StayableStore, error) // bookingIDからstayableなstoreを取得
	GetStoreByQrCode(hash uuid.UUID) (*entity.Store, error)
	GetStoreByUnlimitQrCode(hash uuid.UUID) (*entity.Store, error)
}
