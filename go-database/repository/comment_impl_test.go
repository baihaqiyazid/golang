package repository

import (
	"context"
	"fmt"
	db "go-database"
	"go-database/entity"
	"testing"
)
func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email: "tiza@email.com",
		Comment: "Comment by Tiza",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnection())
	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 45)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnection())
	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, v := range result {
		fmt.Println("====================================")
		fmt.Println(v)
	}
}

func TestCommentUpdate(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnection())
	ctx := context.Background()

	comment := "Comment Updated"
	
	result, err := commentRepository.Update(ctx, 44, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result, comment)
}