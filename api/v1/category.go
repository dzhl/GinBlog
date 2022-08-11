package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		model.CreateCategory(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"stauts": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询分类
func GetCategorys(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageNum == 0 {
		pageNum = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	data := model.GetCategorys(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"staus": code,
		"data":  data,
		"msg":   errmsg.GetErrMsg(code),
	})

}

//编辑分类
func EditCategory(c *gin.Context) {
	//model.CheckUser()
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		code = model.EditCategory(id, &data)
	} else {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
