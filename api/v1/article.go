package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code := model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"stauts": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询分类下所有文章
func GetCatArticles(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageNum == 0 {
		pageNum = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	data := model.GetCatArticles(pageSize, pageNum, cid)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

//查询单个文章
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//var data model.Article
	data := model.GetArticleInfo(id)
	var code int
	if data.ID == 0 {
		code = errmsg.ERROR_ART_NOT_EXIST
	} else {
		code = errmsg.SUCCSE
	}
	c.JSON(http.StatusOK, gin.H{
		"staus": code,
		"data":  data,
		"msg":   errmsg.GetErrMsg(code),
	})

}

//查询所有文章
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageNum == 0 {
		pageNum = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	data := model.GetArticles(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"staus": code,
		"data":  data,
		"msg":   errmsg.GetErrMsg(code),
	})

}

//编辑文章
func EditArticle(c *gin.Context) {
	//model.CheckUser()
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code := model.EditArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
