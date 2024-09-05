package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errpkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/error"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/router/api"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/router/dto"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/service"
)

type AuthAPIV1 struct {
	authService service.AuthService
}

func NewAuthAPIV1(
	authService service.AuthService,
) *AuthAPIV1 {
	return &AuthAPIV1{
		authService: authService,
	}
}

// Register
// @Tags Auth
// @Summary Register
// @Produce json
// @Param data body dto.AuthRegisterRequest true "Register"
// @Success 200 {object} dto.AuthRegisterResponse
// @Router /api/v1/auth/register [post]
func (a *AuthAPIV1) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.AuthRegisterRequest{}

		if err := c.ShouldBindJSON(&req); err != nil {
			api.ResponseError(c, http.StatusBadRequest, errpkg.INVALID_PARAMS, nil)
			return
		}

		user, err := a.authService.Register(c.Request.Context(), req.ToModel())

		if err != nil {
			api.ResponseError(c, http.StatusInternalServerError, errpkg.ERROR_AUTH_REGISTER_FAIL, err)
			return
		}

		api.ResponseSuccess(c, http.StatusOK, dto.AuthRegisterResponse{
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
}

// Login
// @Tags Auth
// @Summary Login
// @Produce json
// @Param data body dto.AuthLoginRequest true "Login"
// @Success 200 {object} dto.AuthLoginResponse
// @Router /api/v1/auth/login [post]
func (a *AuthAPIV1) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.AuthLoginRequest{}

		if err := c.ShouldBindJSON(&req); err != nil {
			api.ResponseError(c, http.StatusBadRequest, errpkg.INVALID_PARAMS, nil)
			return
		}

		token, err := a.authService.Login(c.Request.Context(), req.Username, req.Password)

		if err != nil {
			api.ResponseError(c, http.StatusInternalServerError, errpkg.ERROR_AUTH_LOGIN_FAIL, err)
			return
		}

		api.ResponseSuccess(c, http.StatusOK, dto.AuthLoginResponse{Token: *token})
	}
}
