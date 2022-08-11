package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"stauts": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageNum == 0 {
		pageNum = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"staus": code,
		"data":  data,
		"msg":   errmsg.GetErrMsg(code),
	})

}

//编辑用户列
func EditUser(c *gin.Context) {
	//model.CheckUser()
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		code = model.EditUser(id, &data)
	} else {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
