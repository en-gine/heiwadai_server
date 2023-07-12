package types

type PageQuery struct {
	Limit       *int
	CurrentPage *int
	PerPage     *int
}

// クエリOffset用
func (p *PageQuery) Offset() int {
	return (*p.CurrentPage - 1) * *p.Limit
}
