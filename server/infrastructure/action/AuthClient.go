package action

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"server/core/entity"
	"server/core/infra/action"
	"server/core/infra/types"
	"server/infrastructure/env"

	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
)

var _ action.IAuthAction = &AuthClient{}

type AuthClient struct {
	client      *supa.Client
	userType    action.UserType
	redirectUrl string
}

var (
	authURL          = env.GetEnv(env.SupabaseUrl)
	authKey          = env.GetEnv(env.SupabaseKey)
	adminRedirectURL = env.GetEnv(env.AdminAuthRedirectURL)
	userRedirectURL  = env.GetEnv(env.UserAuthRedirectURL)
)

func NewAuthClient(userType action.UserType) *AuthClient {

	client := supa.CreateClient(authURL, authKey)

	var redirectURL string
	switch userType {
	case action.UserTypeAdmin:
		redirectURL = adminRedirectURL
	case action.UserTypeUser:
	default:
		redirectURL = userRedirectURL
	}

	return &AuthClient{
		client:      client,
		userType:    userType,
		redirectUrl: redirectURL,
	}
}

func (au *AuthClient) SignUp(email string, password entity.Password) (*uuid.UUID, error) {
	ctx := context.Background()
	usr, err := au.signUpWithRedirect(ctx, supa.UserCredentials{
		Email:    email,
		Password: password.String(),
		Data: map[string]interface{}{
			"user_type": au.userType.String(),
		},
	})
	if err != nil {
		return nil, errors.New("Error SignUp" + err.Error())
	}
	userID := uuid.MustParse(usr.ID)
	return &userID, nil
}

func (au *AuthClient) SignIn(email string, password entity.Password) (*types.Token, error) {
	ctx := context.Background()

	auth, err := au.client.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password.String(),
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
	// リダイレクトあり
	ctx := context.Background()
	// client := supa.CreateClient(authURL + "", authKey)
	err := au.resetPasswordForEmailWithRedirect(ctx, email)
	if err != nil {
		return err
	}
	return nil
}

func (au *AuthClient) UpdatePassword(password entity.Password, token string) error {
	ctx := context.Background()

	_, err := au.client.Auth.UpdateUser(ctx, token, map[string]interface{}{
		"password": password.String(),
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
	// リダイレクトあり

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

func (au *AuthClient) signUpWithRedirect(ctx context.Context, credentials supa.UserCredentials) (*supa.User, error) {
	reqBody, _ := json.Marshal(credentials)
	reqURL := fmt.Sprintf("%s/%s/signup?redirect_to=%s", au.client.BaseURL, supa.AuthEndpoint, url.QueryEscape(au.redirectUrl))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res := supa.User{}
	if err := au.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ResetPasswordForEmail sends a password recovery link to the given e-mail address.
func (au *AuthClient) resetPasswordForEmailWithRedirect(ctx context.Context, email string) error {
	reqBody, _ := json.Marshal(map[string]string{"email": email})
	reqURL := fmt.Sprintf("%s/%s/recover?redirect_to=%s", au.client.BaseURL, supa.AuthEndpoint, url.QueryEscape(au.redirectUrl))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	if err = au.sendRequest(req, nil); err != nil {
		return err
	}

	return nil
}

func (au *AuthClient) inviteUserByEmail(ctx context.Context, email string) (*supa.User, error) {
	reqBody, _ := json.Marshal(map[string]string{"email": email})
	reqURL := fmt.Sprintf("%s/%s/invite?redirect_to=%s", au.client.BaseURL, supa.AuthEndpoint, url.QueryEscape(au.redirectUrl))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	au.injectAuthorizationHeader(req, authKey)
	req.Header.Set("Content-Type", "application/json")
	res := supa.User{}
	if err := au.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (au *AuthClient) sendRequest(req *http.Request, v interface{}) error {
	var errRes supa.ErrorResponse
	hasCustomError, err := au.sendCustomRequest(req, v, &errRes)

	if err != nil {
		return err
	} else if hasCustomError {
		return &errRes
	}

	return nil
}

func (au *AuthClient) injectAuthorizationHeader(req *http.Request, value string) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", value))
}

func (au *AuthClient) sendCustomRequest(req *http.Request, successValue interface{}, errorValue interface{}) (bool, error) {
	req.Header.Set("apikey", authKey)
	res, err := au.client.HTTPClient.Do(req)
	if err != nil {
		return true, err
	}

	defer res.Body.Close()
	statusOK := res.StatusCode >= http.StatusOK && res.StatusCode < 300
	if !statusOK {
		if err = json.NewDecoder(res.Body).Decode(&errorValue); err == nil {
			return true, nil
		}

		return false, fmt.Errorf("unknown, status code: %d", res.StatusCode)
	} else if res.StatusCode != http.StatusNoContent {
		if err = json.NewDecoder(res.Body).Decode(&successValue); err != nil {
			return false, err
		}
	}

	return false, nil
}
