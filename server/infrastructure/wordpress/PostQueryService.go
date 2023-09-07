package wordpress

import (
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/infrastructure/logger"
	"server/infrastructure/wordpress/api"
	"server/infrastructure/wordpress/types"
	"time"
)

var _ queryservice.IPostQueryService = &PostQueryService{}

type PostQueryService struct {
}

func NewPostQueryService() *PostQueryService {

	return &PostQueryService{}
}

func (pq *PostQueryService) GetByID(id int) (*entity.Post, error) {
	wppost, err := api.GetWPPost(uint(id))
	if err != nil {
		logger.Errorf("Error: %v\n", err)
		return nil, err
	}
	return WPPostToEntity(wppost), nil
}

func (pq *PostQueryService) GetAll() ([]*entity.Post, error) {
	wpposts, err := api.GetWPPosts()
	if err != nil {
		logger.Errorf("Error: %v\n", err)
		return nil, err
	}
	var entities []*entity.Post
	for _, wppost := range *wpposts {
		entities = append(entities, WPPostToEntity(&wppost))
	}
	return entities, nil
}

func WPPostToEntity(wppost *types.WPPost) (entitie *entity.Post) {
	postDate, _ := time.Parse("2006-01-02T15:04:05", wppost.Date)
	entity := &entity.Post{
		ID:       wppost.ID,
		Title:    wppost.Title.Rendered,
		Content:  wppost.Content.Rendered,
		Author:   wppost.Embedded.Author[0].Name,
		PostDate: postDate,
	}
	return entity
}
