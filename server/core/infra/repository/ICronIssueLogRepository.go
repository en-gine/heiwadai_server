package repository

type ICronIssueLogRepository interface {
	Save(tx ITransaction, cronName string, issueCount int, issueYear int, issueMonth int) error
}
