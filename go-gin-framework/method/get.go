package method

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Muhammad Yazid Baihaqi",
		"username": "baihaqi123456",
	})
}

func BookHandler(ctx *gin.Context)  {
	id := ctx.Param("id")
	
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func FindCategoryHandler(ctx *gin.Context)  {
	title := ctx.Query("title")
	
	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func GetPostHandler(ctx *gin.Context)  {
	id := ctx.Param("id")
	author := ctx.Param("author")
	
	ctx.JSON(http.StatusOK, gin.H{
		"id" : id,
		"author": author,
	})
}

func FindPostHandler(ctx *gin.Context)  {
	id := ctx.Query("id")
	author := ctx.Query("author")
	
	ctx.JSON(http.StatusOK, gin.H{
		"id" : id,
		"author": author,
	})
}