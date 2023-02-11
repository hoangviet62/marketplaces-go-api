package internal

import (
	"errors"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"github.com/hoangviet62/marketplaces-go-api/internal/services/jwt"
	kong "github.com/hoangviet62/marketplaces-go-api/internal/services/kong"
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

	if result.Error != nil {
		signInResp.ErrorMessage = errors.New("invalid username or password")
		return signInResp
	}

	if err := helpers.VerifyPassword(user.Password, signInInput.Password); err != nil {
		signInResp.ErrorMessage = errors.New("invalid username or password")
		return signInResp
	}

	jwtData := kong.FetchConsumerJwt(user.Username)
	jwtToken := jwt.Encode(user, jwtData.Data[0].Key, jwtData.Data[0].Secret)

	signInResp.Status = true
	signInResp.AuthToken = jwtToken
	return signInResp
}
