package main

import (
	"context"
	"testing"

	"net/http"
	"net/http/httptest"
	pb "server/api/v1/user" // your generated code
	userv1connect "server/api/v1/user/userconnect"
	"server/core/usecase/user"

	controller "server/controller/user"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
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

	client := userv1connect.NewAuthControllerClient(
		server.Client(),
		server.URL,
	)

	// grpcの場合
	// grpcClient := userv1connect.NewAuthControllerClient(
	// 	server.Client(),
	// 	server.URL,
	// 	connect.WithGRPC(),
	// )

	msg := &pb.UserRegisterRequest{
		FirstName:     "John",
		LastName:      "Doe",
		FirstNameKana: "ジョン",
		LastNameKana:  "ドウ",
		CompanyName:   "Test Company",
		BirthDate:     "1990-01-01",
		ZipCode:       "1234567",
		Prefecture:    "Tokyo",
		City:          "Shibuya",
		Address:       "1-1-1",
		Tel:           "03-1234-5678",
		Mail:          "test@example.com",
		AcceptMail:    true,
		AcceptTerm:    true,
	}

	t.Run("register", func(t *testing.T) {
		req := &connect.Request[pb.UserRegisterRequest]{Msg: msg}
		res, err := client.Register(context.Background(), req)
		t.Log(res, err)
		assert.Equal(t, res.Msg, emptypb.Empty{})
	})

	t.Run("duplicate register", func(t *testing.T) {
		req := &connect.Request[pb.UserRegisterRequest]{Msg: msg}
		res, err := client.Register(context.Background(), req)
		t.Log(res, err)
		assert.NotNil(t, err)
	})

	t.Run("try sign up", func(t *testing.T) {
		res, err := client.SignUp(context.Background(), &connect.Request[pb.UserAuthRequest]{Msg: &pb.UserAuthRequest{
			Email:    "test@example.com",
			Password: "password",
		}})
		t.Log(res, err)
	})

}
