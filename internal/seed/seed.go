package seed

import (
	"github.com/tranvux/learn-structs/internal/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	// 1. create users
	users := []model.User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
		{Name: "tranvux", Email: "tranvux@example.com"},
		{Name: "baohuy", Email: "baohuy@example.com"},
	}
	// db.Create(&users)
	for _, u := range users {
		db.Where(model.User{Email: u.Email}).FirstOrCreate(&u)
	}

	// 2. create tag
	tags := []model.Tag{
		{Name: "golang"},
		{Name: "nextjs"},
		{Name: "react"},
		{Name: "docker"},
	}
	// db.Create(&tags)
	for _, t := range tags {
		db.Where(model.Tag{Name: t.Name}).FirstOrCreate(&t)
	}

	// 3. create post & assign tag (many2many)
	post := model.Post{
		UserID:  users[2].ID,
		Title:   "Hello gorm",
		Content: "Learning GORM w postgres",
		Tags:    []model.Tag{tags[1], tags[3]},
	}
	// db.Create(&post)
	db.Where(model.Post{Title: post.Title}).FirstOrCreate(&post)

	// 4. create comment
	comment := model.Comment{
		PostID: post.ID,
		UserID: users[1].ID,
		Body:   "Great post!",
	}
	// db.Create(&comment)
	db.Where(model.Comment{PostID: comment.PostID}).FirstOrCreate(&comment)
}
