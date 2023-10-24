package types

const (
	DefaultPerPage     = 10
	DefaultCurrentPage = 1
)

type PageQuery struct {
	CurrentPage *int
	PerPage     *int
}

func NewPageQuery(currentPage *int, perPage *int) *PageQuery {
	if currentPage == nil {
		tmp := DefaultCurrentPage
		currentPage = &tmp
	}
	if perPage == nil {
		tmp := DefaultPerPage
		perPage = &tmp
	}
	return &PageQuery{
		CurrentPage: currentPage,
		PerPage:     perPage,
	}
}

// クエリOffset用
func (p *PageQuery) Offset() int {
	if p.CurrentPage == nil || p.PerPage == nil {
		return 0
	}
	return (*p.CurrentPage - 1) * *p.PerPage
}

func (p *PageQuery) Limit() int {
	if p.PerPage == nil {
		return DefaultPerPage
	}
	return *p.PerPage
}
