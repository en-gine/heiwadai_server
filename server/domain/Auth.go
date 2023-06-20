package domain

type Auth struct {
	email    string
	password string
}

func (a *Auth) New(
	email string,
	password string,
) *Auth {
	return &Auth{
		email:    email,
		password: password,
	}
}
