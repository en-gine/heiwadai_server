package spabase

import (
	"context"
	"errors"
	"os"
	"server/core/entity"
	"server/core/infra/action"
	"server/core/infra/types"
	"server/external/parser"

	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
)

var _ action.IAuthAction = &AuthClient{}

type AuthClient struct {
	client *supa.Client
}

func NewAuthClient() *AuthClient {
	auth := &AuthClient{}
	auth.createClient()
	return &AuthClient{}
}

func (au *AuthClient) createClient() {
	authURL := os.Getenv("SUPABASE_URL")
	authKey := os.Getenv("SUPABASE_KEY")
	if authURL == "" {
		panic("SUPABASE_URL is not set")
	}
	if authKey == "" {
		panic("SUPABASE_KEY is not set")
	}
	au.client = supa.CreateClient(authURL, authKey)

}
func (au *AuthClient) SignUp(email string, password string, UserData *entity.User) (*entity.User, error) {

	ctx := context.Background()
	data, err := au.client.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
		Data:     UserData,
	})
	if err != nil {
		return nil, errors.New("Error SignUp" + err.Error())
	}

	return spaUserToEntity(data)
}

func (au *AuthClient) SignIn(email string, password string) (*types.Token, error) {

	ctx := context.Background()
	auth, err := au.client.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	return &types.Token{
		AccessToken:  auth.AccessToken,
		RefreshToken: &auth.RefreshToken,
		ExpiresIn:    &auth.ExpiresIn,
	}, err

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

func (au *AuthClient) GetUser(token *types.Token) (*entity.User, error) {

	ctx := context.Background()
	data, err := au.client.Auth.User(ctx, token.AccessToken)
	if err != nil {
		return nil, errors.New("Error Getting User by token" + err.Error())
	}

	return spaUserToEntity(data)

}

func (au *AuthClient) ResetPassword(email string) error {
	ctx := context.Background()
	err := au.client.Auth.ResetPasswordForEmail(ctx, email)
	return err
}

// supabaseのユーザー情報をentity.Userに変換
func spaUserToEntity(data *supa.User) (*entity.User, error) {
	uuid, err := uuid.Parse(data.ID)
	if err != nil {
		return nil, errors.New("Error parsing UUID" + err.Error())
	}

	if err != nil {
		return nil, err
	}

	birthDate, err := parser.ToDate(data.UserMetadata["BirthDate"].(string))
	if err != nil {
		return nil, err
	}
	pref, domainErr := entity.IntToPrefecture(data.UserMetadata["Prefecture"].(int))
	if domainErr != nil {
		return nil, domainErr
	}

	return entity.RegenUser(
		uuid,
		data.UserMetadata["FirstName"].(string),
		data.UserMetadata["LastName"].(string),
		data.UserMetadata["FirstNameKana"].(string),
		data.UserMetadata["LastNameKana"].(string),
		parser.ToStringPtr(data.UserMetadata["CompanyName"]),
		birthDate,
		parser.ToStringPtr(data.UserMetadata["ZipCode"]),
		pref,
		parser.ToStringPtr(data.UserMetadata["City"]),
		parser.ToStringPtr(data.UserMetadata["Address"]),
		parser.ToStringPtr(data.UserMetadata["Tel"]),
		data.Email,
		data.UserMetadata["AcceptMail"].(bool),
	), nil

}
