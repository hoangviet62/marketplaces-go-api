package internal

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	kong "github.com/hoangviet62/marketplaces-go-api/internal/services/kong"
)

func CreateUser(context *gin.Context) (bool, model.User, error) {
	// Validate input
	var input model.SignUpInput
	user := model.User{}
	if err := context.ShouldBindJSON(&input); err != nil {
		return false, user, errors.New(err.Error())
	}

	if input.Password != input.PasswordConfirm {
		return false, user, errors.New("password does not match")
	}

	hashedPassword, err := helpers.HashPassword(input.Password)

	if err != nil {
		return false, user, errors.New(err.Error())
	}

	user = model.User{
		Username: strings.ToLower(input.Username),
		Email:    strings.ToLower(input.Email),
		Password: hashedPassword,
		Mobile:   input.Mobile,
		Status:   "inactive",
		// Role:     input.Role,
	}

	result := helpers.DB.Create(&user)

	if result.Error != nil {
		return false, user, errors.New(result.Error.Error())
	}

	status, _, err := kong.CreateConsumer(input.Username)

	if !status {
		return false, user, errors.New(err.Error())
	}

	return status, user, nil
}
