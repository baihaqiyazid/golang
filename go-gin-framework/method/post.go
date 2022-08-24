package method

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct{
	Id json.Number `json:"id" binding:"required,number"`
	Author string `json:"author" binding:"required,alpha"`
	Title string `json:"title" binding:"required"`
}

func CreatePostHandler(ctx *gin.Context)  {
	var post Post
	
	err := ctx.ShouldBindJSON(&post)
	if err != nil {

		var error_request []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error field %s, Condition %s", e.Field(), e.ActualTag())
			error_request = append(error_request, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : error_request,
		})

		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id" : post.Id,
		"author": post.Author, 
		"title": post.Title,
	})
}