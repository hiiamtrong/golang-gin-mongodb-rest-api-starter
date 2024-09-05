package middleware

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	constantpkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/constant"
	errpkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/error"
	jwtpkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/jwt"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/router/api"
)

func (m *AppMiddleware) JWT() gin.HandlerFunc {

	secret := m.config.JWT.Secret

	return func(c *gin.Context) {
		code := errpkg.SUCCESS
		token := c.GetHeader(constantpkg.AuthorizationHeader)
		if token == "" {
			code = errpkg.INVALID_PARAMS
		} else {
			claims, err := jwtpkg.ParseToken(token, secret)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errpkg.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = errpkg.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}

			c.Set(constantpkg.CtxUserClaims, claims)
		}

		if code != errpkg.SUCCESS {
			api.ResponseError(c, http.StatusUnauthorized, code, errors.New(errpkg.GetMsg(code)))
			return
		}

		c.Next()
	}
}
