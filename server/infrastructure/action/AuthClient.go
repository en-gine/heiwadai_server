package action

import (
	"context"
	"errors"
	"os"
	"server/core/infra/action"
	"server/core/infra/types"

	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
)

var _ action.IAuthAction = &AuthClient{}

type AuthClient struct {
	client *supa.Client
}

func NewAuthClient() *AuthClient {
	authURL := os.Getenv("SUPABASE_URL")
	authKey := os.Getenv("SUPABASE_KEY")
	if authURL == "" {
		panic("SUPABASE_URL is not set")
	}
	if authKey == "" {
		panic("SUPABASE_KEY is not set")
	}
	client := supa.CreateClient(authURL, authKey)

	return &AuthClient{
		client: client,
	}
}

func (au *AuthClient) SignUp(email string, password string, userType action.UserType) error {

	ctx := context.Background()
	_, err := au.client.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
		Data:     map[string]string{"userType": userType.String()},
	})
	if err != nil {
		return errors.New("Error SignUp" + err.Error())
	}

	return nil
}

func (au *AuthClient) SignIn(email string, password string) (*types.Token, error) {

	ctx := context.Background()
	auth, err := au.client.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, errors.New("Error SignIn" + err.Error())
	}

	return &types.Token{
		AccessToken:  auth.AccessToken,
		RefreshToken: &auth.RefreshToken,
		ExpiresIn:    &auth.ExpiresIn,
	}, nil
}

func (au *AuthClient) Refresh(token *types.Token) (*types.Token, error) {

	ctx := context.Background()
	data, err := au.client.Auth.RefreshUser(ctx, token.AccessToken, *token.RefreshToken)
	if err != nil {
		return nil, errors.New("Error Refreshing Token" + err.Error())
	}

	return &types.Token{
		AccessToken:  data.AccessToken,
		RefreshToken: &data.RefreshToken,
		ExpiresIn:    &data.ExpiresIn,
	}, err
}

func (au *AuthClient) ResetPasswordMail(email string) error {
	ctx := context.Background()
	err := au.client.Auth.ResetPasswordForEmail(ctx, email)
	return err
}

func (au *AuthClient) UpdatePassword(password string, token string) error {
	ctx := context.Background()
	_, err := au.client.Auth.UpdateUser(ctx, token, map[string]interface{}{
		"password": password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (au *AuthClient) InviteUserByEmail(email string) (uuid.UUID, error) {
	ctx := context.Background()
	user, err := au.client.Auth.InviteUserByEmail(ctx, email)
	if err != nil {
		return uuid.Nil, err
	}
	newUserID := uuid.MustParse(user.ID)
	return newUserID, nil
}

func (au *AuthClient) UpdateEmail(email string, token string) error {
	ctx := context.Background()
	_, err := au.client.Auth.UpdateUser(ctx, token, map[string]interface{}{
		"email": email,
	})
	return err
}
