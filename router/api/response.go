package api

import (
	"github.com/gin-gonic/gin"
	errpkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/error"
)

type response struct {
	Code errpkg.Code `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, response{
		Code: errpkg.SUCCESS,
		Msg:  errpkg.GetMsg(errpkg.SUCCESS),
		Data: data,
	})
	return
}

func ResponseError(c *gin.Context, httpCode int, errCode errpkg.Code, err error) {
	c.AbortWithStatusJSON(httpCode, response{
		Code: errCode,
		Msg:  errpkg.GetMsg(errCode),
		Data: err.Error(),
	})
	return
}
