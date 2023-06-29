package types

type Token struct {
	AccessToken  string
	ExpiresIn    *int
	RefreshToken *string
}
