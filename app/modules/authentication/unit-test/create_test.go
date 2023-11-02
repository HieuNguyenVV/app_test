package unittest

import (
	"app/common"
	"app/modules/authentication/mock"
	"app/modules/authentication/model"
	"app/modules/authentication/service"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	createRepo := mock.NewMockCreateAccountRepository(mockCtrl)
	createService := service.NewCreateAccountService(createRepo)

	type (
		testCaseInput struct {
			request model.CreateAccountRequest
		}
		testCaseMockedFn struct {
			mockData func(model.CreateAccountRequest)
		}
		testCaseOutput struct {
			err error
		}
	)
	var (
		errDB    = errors.New("something went wrong with db")
		msg      = "email existed"
		errExist = errors.New(msg)
		//errHasedPass = errors.New("hashed password fail")
	)
	listTestCase := []struct {
		name         string
		input        testCaseInput
		fn           testCaseMockedFn
		expectOutput testCaseOutput
	}{
		{
			name: "1: Error when check exist email",
			input: testCaseInput{
				request: model.CreateAccountRequest{
					Username: "Nguyen Van Hieu",
					Password: "1234567890",
					Email:    "hieunguyenvan989@gmail.com.vn",
				},
			},
			fn: testCaseMockedFn{
				mockData: func(model.CreateAccountRequest) {
					createRepo.EXPECT().CheckExistByEmail(context.Background(), gomock.Any()).Return(false, common.ErrorDB(errDB))
					//createRepo.EXPECT().CreateAccount(context.Background(), gomock.Any()).Return(nil)
				},
			},
			expectOutput: testCaseOutput{
				err: common.ErrorDB(errDB),
			},
		},
		{
			name: "2: Account existed",
			input: testCaseInput{
				request: model.CreateAccountRequest{
					Username: "Nguyen Van Hieu",
					Password: "1234567890",
					Email:    "hieunguyenvan989@gmail.com.vn",
				},
			},
			fn: testCaseMockedFn{
				mockData: func(model.CreateAccountRequest) {
					createRepo.EXPECT().CheckExistByEmail(context.Background(), gomock.Any()).Return(true, nil)
				},
			},
			expectOutput: testCaseOutput{
				err: common.NewErrorResponse(errExist, msg, errExist.Error(), "EMAIL_EXISTED"),
			},
		},
		{
			name: "4: Create account fail",
			input: testCaseInput{
				request: model.CreateAccountRequest{
					Username: "Nguyen Van Hieu",
					Password: "1234567890",
					Email:    "hieunguyenvan989@gmail.com.vn",
				},
			},
			fn: testCaseMockedFn{
				mockData: func(model.CreateAccountRequest) {
					createRepo.EXPECT().CheckExistByEmail(context.Background(), gomock.Any()).Return(false, nil)
					createRepo.EXPECT().CreateAccount(context.Background(), gomock.Any()).Return(common.ErrorDB(errDB))
				},
			},
			expectOutput: testCaseOutput{
				err: common.ErrorDB(errDB),
			},
		},
		{
			name: "5: Success",
			input: testCaseInput{
				request: model.CreateAccountRequest{
					Username: "Nguyen Van Hieu",
					Password: "1234567890",
					Email:    "hieunguyenvan989@gmail.com.vn",
				},
			},
			fn: testCaseMockedFn{
				mockData: func(model.CreateAccountRequest) {
					createRepo.EXPECT().CheckExistByEmail(context.Background(), gomock.Any()).Return(false, nil)
					createRepo.EXPECT().CreateAccount(context.Background(), gomock.Any()).Return(nil)
				},
			},
			expectOutput: testCaseOutput{
				err: nil,
			},
		},
	}

	for k, tc := range listTestCase {
		t.Run(fmt.Sprintf("%d. %s", k, tc.name), func(t *testing.T) {
			tc.fn.mockData(tc.input.request)
			err := createService.CreateAccount(context.Background(), &tc.input.request)
			if err == nil && tc.expectOutput.err == nil {
				return
			}
			if err.Error() != tc.expectOutput.err.Error() {
				t.Errorf("Test create account fail, exp: %s, got: %v", tc.expectOutput.err.Error(), err.Error())
			}
		})
	}
}
