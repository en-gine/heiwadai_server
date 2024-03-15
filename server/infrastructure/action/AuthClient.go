package action

import (
	"context"
	"errors"

	"server/core/infra/action"
	"server/core/infra/types"
	"server/infrastructure/env"

	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
)

var _ action.IAuthAction = &AuthClient{}

type AuthClient struct {
	client *supa.Client
	apiKey string
}

func NewAuthClient() *AuthClient {
	authURL := env.GetEnv(env.SupabaseUrl)
	authKey := env.GetEnv(env.SupabaseKey)
	if authURL == "" {
		panic("SUPABASE_URL is not set")
	}
	if authKey == "" {
		panic("SUPABASE_KEY is not set")
	}
	client := supa.CreateClient(authURL, authKey)

	return &AuthClient{
		client: client,
		apiKey: authKey,
	}
}

func (au *AuthClient) SignUp(email string, password string, userType action.UserType) (*uuid.UUID, error) {
	ctx := context.Background()
	usr, err := au.client.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
		Data: map[string]interface{}{
			"user_type": userType.String(),
		},
	})
	if err != nil {
		return nil, errors.New("Error SignUp" + err.Error())
	}
	userID := uuid.MustParse(usr.ID)
	return &userID, nil
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

func (au *AuthClient) SignOut(token string) error {
	ctx := context.Background()
	err := au.client.Auth.SignOut(ctx, token)
	if err != nil {
		return errors.New("Error SignIn" + err.Error())
	}

	return nil
}

func (au *AuthClient) Refresh(token string, refreshToken string) (*action.UserAuth, error) {
	ctx := context.Background()
	data, err := au.client.Auth.RefreshUser(ctx, token, refreshToken)
	if err != nil {
		return nil, errors.New("Error Refreshing Token" + err.Error())
	}
	Token := &types.Token{
		AccessToken:  data.AccessToken,
		RefreshToken: &data.RefreshToken,
		ExpiresIn:    &data.ExpiresIn,
	}
	ut := data.User.UserMetadata["user_type"]
	var userType string
	if ut == "" || ut == nil {
		userType = "user"
	} else {
		userType = ut.(string)
	}
	return &action.UserAuth{
		UserID:   uuid.MustParse(data.User.ID),
		UserType: action.UserType(userType),
		Token:    Token,
	}, nil
}

func (au *AuthClient) ResetPasswordMail(email string) error {
	ctx := context.Background()
	err := au.client.Auth.ResetPasswordForEmail(ctx, email)
	if err != nil {
		return err
	}
	return nil
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

func (au *AuthClient) InviteUserByEmail(email string) (*uuid.UUID, error) {
	ctx := context.Background()
	user, err := au.client.Auth.InviteUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	newUserID := uuid.MustParse(user.ID)
	return &newUserID, nil
}

func (au *AuthClient) UpdateEmail(email string, token string) error {
	ctx := context.Background()
	_, err := au.client.Auth.UpdateUser(ctx, token, map[string]interface{}{
		"email": email,
	})
	if err != nil {
		return err
	}
	return nil
}

func (au *AuthClient) GetUserID(token string) (*uuid.UUID, *action.UserType, error) {
	ctx := context.Background()

	user, err := au.client.Auth.User(ctx, token)
	if err != nil {
		return nil, nil, err
	}
	userID := uuid.MustParse(user.ID)
	metadata := user.UserMetadata["user_type"].(string)
	userType := action.UserType(metadata)
	return &userID, &userType, nil
}
