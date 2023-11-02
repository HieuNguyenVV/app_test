package unit_test

import (
	"app/modules/authentication/controller"
	"app/modules/authentication/model"
	"app/pkg"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	app := pkg.App()

	r := gin.Default()
	r.POST("/create", controller.CreateAccount(app))

	type (
		testCaseInput model.CreateAccountRequest
	)
	testCases := []struct {
		name  string
		input testCaseInput
		code  int
	}{
		{
			name: "1: Something went wrong with database, should return 400",
			input: testCaseInput{
				Username: "Nguyen Van A",
				Password: "1234567890",
				Email:    "nguyenvanA@gmail.com.vn",
			},
			code: 400,
		},
		{
			name: "2: Error validation, should return 400",
			input: testCaseInput{
				Username: "Nguyen Van A",
				Password: "",
				Email:    "nguyenvanABC@gmail.com.vn",
			},
			code: 400,
		},
		// {
		// 	name: "3: Something went wrong with server, should return 500",
		// 	input: testCaseInput{
		// 		Username: "Nguyen Van A",
		// 		Password: "1234567890",
		// 		Email:    "nguyenvanA@gmail.com.vn",
		// 	},
		// 	code: 500,
		// },
		{
			name: "3: Success",
			input: testCaseInput{
				Username: "Nguyen Van D",
				Password: "1234567890",
				Email:    "nguyenvanABCD@gmail.com.vn",
			},
			code: 200,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			jsonInput, err := json.Marshal(&v.input)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(jsonInput))
			if err != nil {
				t.Fatalf("Error: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			if ok := assert.Equal(t, v.code, res.Code); !ok {
				t.Fatalf("Test fail, exp: %v, got: %v", v.code, res.Code)
			}
		})
	}

}

func TestLogin(t *testing.T) {
	app := pkg.App()
	r := gin.Default()
	r.POST("/login", controller.Login(app))

	testCases := []struct {
		name  string
		input model.AccountLoginRequest
		resp  model.AccountLoginResponse
		code  int
	}{
		{
			name: "1: Invalid request, should return 400",
			input: model.AccountLoginRequest{
				Password: "",
				Email:    "nguyenvanABCD@gmail.com.vn",
			},
			resp: model.AccountLoginResponse{},
			code: 400,
		},
		{
			name: "2: Something went wrong with database, should return 400",
			input: model.AccountLoginRequest{
				Password: "1234567890",
				Email:    "nguyenvanABCDE@gmail.com.vn",
			},
			resp: model.AccountLoginResponse{},
			code: 400,
		},
		{
			name: "3: Something went wrong with server, should return 500",
			input: model.AccountLoginRequest{
				Password: "1234567890123",
				Email:    "nguyenvanABCD@gmail.com.vn",
			},
			resp: model.AccountLoginResponse{},
			code: 500,
		},
		{
			name: "4: Success, should return 200",
			input: model.AccountLoginRequest{
				Password: "1234567890",
				Email:    "nguyenvanABCD@gmail.com.vn",
			},
			resp: model.AccountLoginResponse{},
			code: 200,
		},
	}
	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			jsonInput, err := json.Marshal(&v.input)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonInput))
			if err != nil {
				t.Fatalf("Error: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			if ok := assert.Equal(t, v.code, res.Code); !ok {
				t.Fatalf("Test fail, exp: %v, got: %v", v.code, res.Code)
			}
		})
	}
}

func TestChangePassword(t *testing.T) {
	app := pkg.App()
	r := gin.Default()

	r.POST("/change", controller.ChangePassword(app))
	input := model.ChangePasswordRequest{
		OldPassword:     "1234567890",
		NewPassword:     "1234567890",
		ConfirmPassword: "1234567890",
	}
	inputJson, err := json.Marshal(&input)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "/change", bytes.NewBuffer(inputJson))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjRmMDJjMzIyLWQwYTYtMTFlZC04OTVkLTMyMDNjODIyNDBhMSIsImFjY291bnRJZCI6OSwicm9sZSI6MSwiZXhwaXJlZEF0IjoiMjAyMy0wNC0wMlQxOTowMDoyMi43MTIzNzg5KzA3OjAwIiwiaXNzdWVBdCI6IjIwMjMtMDQtMDFUMjM6MDA6MjIuNzEyMzc4OSswNzowMCJ9.PHGRn4r-wiiwQ2-klbGvfl84CA44OpqLRIcWjf6mjI4")

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if ok := assert.Equal(t, http.StatusOK, res.Code); ok {
		t.Fatalf("Test fail, exp: %v, got: %v", http.StatusOK, res.Code)
	}
}
