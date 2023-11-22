package types

type PageResponse struct {
	CurrentPage int
	PerPage     int
	TotalCount  int
	TotalPage   int
}

func NewPageResponse(pager *PageQuery, totalCount int) *PageResponse {
	var totalPage int
	if pager.PerPage == nil || *pager.PerPage == 0 {
		totalPage = 1
	} else {
		totalPage = (totalCount + *pager.PerPage - 1) / *pager.PerPage
	}
	return &PageResponse{
		CurrentPage: *pager.CurrentPage,
		PerPage:     *pager.PerPage,
		TotalCount:  totalCount,
		TotalPage:   totalPage,
	}
}
