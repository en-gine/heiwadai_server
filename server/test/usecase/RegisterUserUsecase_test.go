package usecase_test

// Path: test/usecase/RegisterUserUsecaseTest.go
// Compare this snippet from usecase/user/RegisterUser.go:
// package user
//
import (
	"reflect"
	"server/domain"
	"server/infra/interfaces"
	"server/usecase/user"
	"testing"
	"time"

	"errors"
	"server/infra/mocks"

	"github.com/google/uuid"
)

type RegisterUserUsecase struct {
	userRepository interfaces.IUserRepository
}

func Test_RegisterUserUsecase_Exec(t *testing.T) {
	// Create a mock user repository
	mockRepo := new(mocks.UserRepository)

	// Create a RegisterUserUsecase instance with the mock repository
	usecase := user.New(mockRepo)

	// Define test cases
	tests := []struct {
		name           string
		firstName      string
		lastName       string
		firstNameKana  string
		lastNameKana   string
		companyName    *string
		birthDate      time.Time
		zipCode        *string
		prefecture     string
		city           *string
		address        *string
		tel            *string
		mail           string
		acceptMail     bool
		existingUser   *domain.User
		findByMailErr  error
		expectedError  error
		expectedResult *domain.User
	}{
		{
			name:          "Successfully register new user",
			firstName:     "John",
			lastName:      "Doe",
			firstNameKana: "ジョン",
			lastNameKana:  "ドウ",
			birthDate:     time.Now(),
			prefecture:    "Tokyo",
			mail:          "john.doe@example.com",
			acceptMail:    true,
			existingUser:  nil,
			findByMailErr: nil,
			expectedError: nil,
			expectedResult: &domain.User{
				UUID:          uuid.New(),
				FirstName:     "John",
				LastName:      "Doe",
				FirstNameKana: "ジョン",
				LastNameKana:  "ドウ",
				BirthDate:     time.Now(),
				Prefecture:    "Tokyo",
				Mail:          "john.doe@example.com",
				AcceptMail:    true,
			},
		},
		{
			name:           "Fail to register user with existing email",
			firstName:      "John",
			lastName:       "Doe",
			firstNameKana:  "ジョン",
			lastNameKana:   "ドウ",
			birthDate:      time.Now(),
			prefecture:     "Tokyo",
			mail:           "john.doe@example.com",
			acceptMail:     true,
			existingUser:   &domain.User{},
			findByMailErr:  nil,
			expectedError:  errors.New("既に登録されているメールアドレスです"),
			expectedResult: nil,
		},
		{
			name:          "Fail to register user when repository search fails",
			firstName:     "John",
			lastName:      "Doe",
			firstNameKana: "ジョン",
			lastNameKana:  "ドウ",
			birthDate:     time.Now(),
			prefecture:    "Tokyo",
			mail:          "john.doe@example.com",
			acceptMail:    true,
			existingUser:  nil,
			findByMailErr: errors.New("error searching for user"),
			expectedError: errors.New("ユーザーの検索に失敗しました"),
		},
	}

	// Run the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock repository to return the existing user and error as specified in the test case
			mockRepo.On("FindByMail", tt.mail).Return(tt.existingUser, tt.findByMailErr)

			// Call the Exec method on the usecase with the input parameters from the test case
			result, err := usecase.Exec(
				tt.firstName,
				tt.lastName,
				tt.firstNameKana,
				tt.lastNameKana,
				tt.companyName,
				tt.birthDate,
				tt.zipCode,
				tt.prefecture,
				tt.city,
				tt.address,
				tt.tel,
				tt.mail,
				tt.acceptMail,
			)

			// Check that the returned result and error match the expected values from the test case
			if (err != nil && tt.expectedError == nil) || (err != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("Expected error '%v' but got '%v'", tt.expectedError, err)
			}
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("Expected result '%v' but got '%v'", tt.expectedResult, result)
			}

			// Assert that the FindByMail method on the mock repository was called once with the correct parameters
			mockRepo.AssertCalled(t, "FindByMail", tt.mail)
		})
	}
}
