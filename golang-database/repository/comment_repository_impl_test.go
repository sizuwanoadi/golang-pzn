package repository

import (
	"context"
	"fmt"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "adi15siswanto@gmail.com",
		Comment: "Halo Bro, Adi",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindByID(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment, err := CommentRepository.FindById(ctx, 5)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comments, err := CommentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(comments)
}
