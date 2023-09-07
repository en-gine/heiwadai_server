package user

import (
	"context"
	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	usecase "server/core/usecase/user"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostController struct {
	postUseCase usecase.PostUsecase
}

var _ userv1connect.PostControllerClient = &PostController{}

func NewPostController(postUsecase *usecase.PostUsecase) *PostController {
	return &PostController{
		postUseCase: *postUsecase,
	}
}
func (ac *PostController) GetPostByID(ctx context.Context, req *connect.Request[user.PostRequest]) (*connect.Response[user.PostResponse], error) {
	msg := req.Msg
	postID := msg.ID

	post, domaiErr := ac.postUseCase.GetByID(postID)
	return connect.NewResponse(&user.PostResponse{
		ID:       uint32(post.ID),
		Title:    post.Title,
		Content:  post.Content,
		Author:   post.Author,
		PostDate: post.PostDate.String(),
	}), domaiErr
}

func (ac *PostController) GetPosts(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.PostsResponse], error) {
	posts, domaiErr := ac.postUseCase.GetList()

	var resPosts []*user.PostResponse
	for _, post := range posts {
		resPosts = append(resPosts, &user.PostResponse{
			ID:       uint32(post.ID),
			Title:    post.Title,
			Content:  post.Content,
			Author:   post.Author,
			PostDate: post.PostDate.String(),
		})
	}

	return connect.NewResponse(&user.PostsResponse{
		Posts: resPosts,
	}), domaiErr
}
