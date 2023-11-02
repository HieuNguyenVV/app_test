package unittest

import (
	"app/common"
	"app/modules/authentication/mock"
	"app/modules/authentication/model"
	"app/modules/authentication/service"
	"app/pkg/token"
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repoLogin := mock.NewMockLoginRepository(mockCtrl)
	test, _ := time.ParseDuration("10h")
	tokenMaker, _ := token.NewJWTMaker("")
	serviceLogin := service.NewLoginService(repoLogin, tokenMaker, test, test)

	type (
		testCaseInput struct {
			request model.AccountLoginRequest
		}
		testCaseFn struct {
			fn func(model.AccountLoginRequest)
		}
		testCaseOutput struct {
			resp *model.AccountLoginResponse
			err  error
		}
	)
	var (
		errDB    = errors.New("something went wrong with db")
		msg      = "email does not exist"
		errExist = errors.New(msg)
		//errHasedPass = errors.New("hashed password fail")
		errCheckFail = errors.New("crypto/bcrypt: hashedSecret too short to be a bcrypted password")
	)
	listCase := []struct {
		name    string
		input   testCaseInput
		mockFun testCaseFn
		output  testCaseOutput
	}{
		{
			name: "1: Error when check exist email",
			input: testCaseInput{
				request: model.AccountLoginRequest{
					Email:    "hieunguyenvan989@gmail.com.vn",
					Password: "1234567890",
				},
			},
			mockFun: testCaseFn{
				fn: func(model.AccountLoginRequest) {
					repoLogin.EXPECT().CheckExistByEmail(context.Background(), gomock.Any()).Return(false, common.ErrorDB(errDB))
				},
			},
			output: testCaseOutput{
				resp: &model.AccountLoginResponse{
					Account:               model.AccountLoginEntity{},
					AccessToken:           "",
					AccessTokenExpiredAt:  time.Time{},
					RefreshToken:          "",
					RefreshTokenExpiredAt: time.Time{},
					LastLoginAt:           time.Time{},
				},
				err: common.ErrorDB(errDB),
			},
		},

		{
			name: "2: Account not exist",
			input: testCaseInput{
				request: model.AccountLoginRequest{
					Email:    "hieunguyenvan989@gmail.com.vn",
					Password: "1234567890",
				},
			},
			mockFun: testCaseFn{
				fn: func(model.AccountLoginRequest) {
					repoLogin.EXPECT().CheckExistByEmail(context.Background(), gomock.Any()).Return(false, nil)
				},
			},
			output: testCaseOutput{
				resp: &model.AccountLoginResponse{
					Account:               model.AccountLoginEntity{},
					AccessToken:           "",
					AccessTokenExpiredAt:  time.Time{},
					RefreshToken:          "",
					RefreshTokenExpiredAt: time.Time{},
					LastLoginAt:           time.Time{},
				},
				err: common.NewErrorResponse(errExist, msg, errExist.Error(), "EMAIL_NOT_EXIST"),
			},
		},

		{
			name: "3: Error get account",
			input: testCaseInput{
				request: model.AccountLoginRequest{
					Email:    "hieunguyenvan989@gmail.com.vn",
					Password: "1234567890",
				},
			},
			mockFun: testCaseFn{
				fn: func(model.AccountLoginRequest) {
					repoLogin.EXPECT().CheckExistByEmail(context.Background(), gomock.Any()).Return(true, nil)
					repoLogin.EXPECT().GetAccountByEmail(context.Background(), gomock.Any()).Return(&model.Account{}, common.ErrorDB(errDB))
				},
			},
			output: testCaseOutput{
				resp: &model.AccountLoginResponse{
					Account:               model.AccountLoginEntity{},
					AccessToken:           "",
					AccessTokenExpiredAt:  time.Time{},
					RefreshToken:          "",
					RefreshTokenExpiredAt: time.Time{},
					LastLoginAt:           time.Time{},
				},
				err: common.ErrorDB(errDB),
			},
		},

		{
			name: "4: Check pass fail",
			input: testCaseInput{
				request: model.AccountLoginRequest{
					Email:    "hieunguyenvan989@gmail.com.vn",
					Password: "1234567890",
				},
			},
			mockFun: testCaseFn{
				fn: func(model.AccountLoginRequest) {
					repoLogin.EXPECT().CheckExistByEmail(context.Background(), gomock.Any()).Return(true, nil)
					repoLogin.EXPECT().GetAccountByEmail(context.Background(), gomock.Any()).Return(&model.Account{Password: "12345"}, nil)
				},
			},
			output: testCaseOutput{
				resp: &model.AccountLoginResponse{
					Account:               model.AccountLoginEntity{},
					AccessToken:           "",
					AccessTokenExpiredAt:  time.Time{},
					RefreshToken:          "",
					RefreshTokenExpiredAt: time.Time{},
					LastLoginAt:           time.Time{},
				},
				err: common.ErrInternal(errCheckFail),
			},
		},
	}

	for k, tc := range listCase {
		t.Run(fmt.Sprintf("%v. %v", k, tc.name), func(t *testing.T) {
			tc.mockFun.fn(tc.input.request)
			resp, err := serviceLogin.Login(context.Background(), &tc.input.request)
			if reflect.DeepEqual(resp, tc.output.resp) {
				t.Errorf("Test login fail, exp: %v, got: %v", tc.output.resp, resp)
			} else {
				if err.Error() != tc.output.err.Error() {
					t.Errorf("Test login fail, exp: %v, got: %v", tc.output.err.Error(), err.Error())
				}
			}
		})
	}

}
