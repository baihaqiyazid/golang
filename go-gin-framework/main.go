package main

import (
	"go-gin-framework/method"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//GET
	router.GET("/", method.RootHandler) //root
	router.GET("book/:id", method.BookHandler) //param
	router.GET("categories/", method.FindCategoryHandler) //query
	router.GET("post/:id/:author", method.GetPostHandler)
	router.GET("post/", method.FindPostHandler)

	//POST
	router.POST("post/", method.CreatePostHandler)

	router.Run()
}
