package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"strings"
)

func CreateUser(context *gin.Context) (bool, error) {
	// Validate input
	var input model.SignUpInput
	if err := context.ShouldBindJSON(&input); err != nil {
		return false, errors.New(err.Error())
	}

	if input.Password != input.PasswordConfirm {
		return false, errors.New("Passwords do not match")
	}

	hashedPassword, err := helpers.HashPassword(input.Password)

	if err != nil {
		return false, errors.New(err.Error())
	}

	user := model.User{
		Username: input.Username,
		Email:    strings.ToLower(input.Email),
		Password: hashedPassword,
	}

	result := helpers.DB.Create(&user)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return false, errors.New("User with that email already exists")
	} else if result.Error != nil {
		return false, errors.New("Something bad happened")
	}

	return true, nil
}
