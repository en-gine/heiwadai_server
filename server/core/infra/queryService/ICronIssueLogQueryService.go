package queryservice

type ICronIssueLogQueryService interface {
	HasYearMonthLog(year int, month int) (bool, error)
}
