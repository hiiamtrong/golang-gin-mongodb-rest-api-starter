package dto

import (
	"time"

	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/model"
)

type AuthRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (req *AuthRegisterRequest) Validate() error {
	return nil
}

func (req *AuthRegisterRequest) ToModel() model.User {
	user := model.NewUser()
	user.Username = req.Username
	user.Password = req.Password
	return user
}

type AuthRegisterResponse struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
}
