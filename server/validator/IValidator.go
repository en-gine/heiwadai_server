package validator

type IValidator interface {
	IsValid(email string) (bool, error)
}
