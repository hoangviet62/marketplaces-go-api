package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
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
	signInInput := SignInInput{}
	signInResp := SignInResponse{Status: false}

	if err := context.ShouldBindJSON(&signInInput); err != nil {
		signInResp.ErrorMessage = errors.New(err.Error())
		return signInResp
	}

	user := model.User{Username: signInInput.Username}
	helpers.DB.Find(&user)
	// match := helpers.CheckPasswordHash(signInInput.Password, user.Password)
	// pp.Print(match)
	// if errorPass != nil {
	// 	signInResp.ErrorMessage = errors.New("password invalid")
	// 	return signInResp
	// }

	signInResp.Status = true
	return signInResp
}
