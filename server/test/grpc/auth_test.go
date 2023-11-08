package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"server/api/v1/shared"
	pb "server/api/v1/user" // your generated code
	userv1connect "server/api/v1/user/userconnect"
	"server/core/usecase/user"

	controller "server/controller/user"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// const (
// 	address = "localhost:3000"
// )

func TestAuthController_Register(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mux.Handle(userv1connect.NewAuthControllerHandler(
		controller.NewAuthController(&user.AuthUsecase{}),
	))
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()

	authClient := userv1connect.NewAuthControllerClient(
		server.Client(),
		server.URL,
	)

	// grpcの場合
	// grpcClient := userv1connect.NewAuthControllerClient(
	// 	server.Client(),
	// 	server.URL,
	// 	connect.WithGRPC(),
	// )
	birth, _ := time.Parse("20060102", "20181220")
	company := "Test Company"
	zipCode := "1234567"
	city := "Shibuya"
	address := "1-1-1"
	tel := "03-1234-5678"

	msg := &pb.UserRegisterRequest{
		FirstName:     "John",
		LastName:      "Doe",
		FirstNameKana: "ジョン",
		LastNameKana:  "ドウ",
		CompanyName:   &company,
		BirthDate:     timestamppb.New(birth),
		ZipCode:       &zipCode,
		Prefecture:    shared.Prefecture(13),
		City:          &city,
		Address:       &address,
		Tel:           &tel,
		Mail:          "test@example.com",
		AcceptMail:    true,
		AcceptTerm:    true,
	}

	t.Run("register", func(t *testing.T) {
		req := &connect.Request[pb.UserRegisterRequest]{Msg: msg}
		res, err := authClient.Register(context.Background(), req)
		t.Log(res, err)
		assert.Equal(t, res.Msg, emptypb.Empty{})
	})

	t.Run("duplicate register", func(t *testing.T) {
		req := &connect.Request[pb.UserRegisterRequest]{Msg: msg}
		res, err := authClient.Register(context.Background(), req)
		t.Log(res, err)
		assert.NotNil(t, err)
	})

	t.Run("try sign up", func(t *testing.T) {
		res, err := authClient.SignUp(context.Background(), &connect.Request[pb.UserAuthRequest]{Msg: &pb.UserAuthRequest{
			Email:    "test@example.com",
			Password: "password",
		}})
		t.Log(res, err)
	})
}
