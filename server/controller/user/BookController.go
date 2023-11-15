package user

import (
	usecase "server/core/usecase/user"
)

type BookController struct {
	bookUseCase usecase.BookUsecase
}

// var _ userv1connect.BookControllerClient = &BookController{}

func NewBookController(bannerUsecase *usecase.BookUsecase) *BookController {
	return &BookController{
		bookUseCase: *bannerUsecase,
	}
}

// func (ac *BookController) GetBook(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.BookResponse], error) {
// 	book, err := ac.bookUseCase.GetMyBook()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return connect.NewResponse(&user.BookResponse{
// 		Book: book,
// 	}), nil
// }
