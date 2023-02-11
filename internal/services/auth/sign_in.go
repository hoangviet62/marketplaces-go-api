package internal

import (
	"errors"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"strings"
)

type SignInResponse struct {
	Status       bool
	AuthToken    string
	ErrorMessage error
}

type SignInInput struct {
	Username string
	Password string
}

func SignIn(context *gin.Context) SignInResponse {
	signInInput := model.SignInInput{}
	signInResp := SignInResponse{Status: false}

	if err := context.ShouldBindJSON(&signInInput); err != nil {
		signInResp.ErrorMessage = errors.New(err.Error())
		return signInResp
	}

	var user model.User
	result := helpers.DB.First(&user, "username = ?", strings.ToLower(signInInput.Username))

	fmt.Println("result", &user.Password)
	if result.Error != nil {
		signInResp.ErrorMessage = errors.New("Invalid email or Password")
		return signInResp
	}

	if err := helpers.VerifyPassword(user.Password, signInInput.Password); err != nil {
		signInResp.ErrorMessage = errors.New("Invalid email or Password")
		return signInResp
	}

	// match := helpers.CheckPasswordHash(signInInput.Password, user.Password)
	// pp.Print(match)
	// if errorPass != nil {
	// 	signInResp.ErrorMessage = errors.New("password invalid")
	// 	return signInResp
	// }

	signInResp.Status = true
	return signInResp
}
