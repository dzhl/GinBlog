package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	code := model.CheckLogin(data.Username, data.Password)
	var tokenStr string
	if code == errmsg.SUCCSE {
		tokenStr, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"token":  tokenStr,
		"msg":    errmsg.GetErrMsg(code),
	})
}
