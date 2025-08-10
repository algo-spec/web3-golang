package sql_gorm

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

type User struct {
	gorm.Model
	Name      string
	Email     string
	Posts     []Post
	PostCount int64
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint
	User          User
	Comments      []Comment
	CommentStatus string
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	UserID  uint
	Post    Post
	User    User
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("评论:", c.ID, "文章:", c.PostID, "用户:", c.UserID)
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			UpdateColumn("comment_status", "无评论").Error
	}
	return nil
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	var users []User
	for i := 1; i <= 5; i++ {
		var posts []Post
		for j := 1; j <= 3; j++ {
			var comments []Comment
			for k := 1; k <= 2; k++ {
				comments = append(comments, Comment{
					Content: "comment-" + strconv.Itoa(k) + "-post-" + strconv.Itoa(j) + "-user-" + strconv.Itoa(i),
					UserID:  uint(i),
					PostID:  uint(j),
				})
			}
			posts = append(posts, Post{
				Title:    "post-" + strconv.Itoa(j) + "-user-" + strconv.Itoa(i),
				Content:  "content-" + strconv.Itoa(j) + "-user-" + strconv.Itoa(i),
				UserID:   uint(i),
				Comments: comments,
			})
		}
		users = append(users, User{
			Name:  "user-" + strconv.Itoa(i),
			Email: "email-" + strconv.Itoa(i),
			Posts: posts,
		})
	}
	db.Create(&users)

	fmt.Println("====================================================================")
	
	var user User
	db.Debug().Preload("Posts.Comments.User").Find(&user, 1)
	for _, post := range user.Posts {
		fmt.Println(user.Name, "发布的文章：", post.Title)
		for _, comment := range post.Comments {
			fmt.Println("评论内容：", comment.Content, "由用户：", comment.User.Name)
		}
	}

	fmt.Println("====================================================================")

	var mostCommentedPost Post
	db.Debug().Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&mostCommentedPost)

	var commentCount int64
	db.Model(&Comment{}).Where("post_id = ?", mostCommentedPost.ID).Count(&commentCount)

	fmt.Println("评论数量最多的文章：", mostCommentedPost.Title, "评论数量：", commentCount)

	fmt.Println("====================================================================")

	var comment Comment
	db.First(&comment, 2)
	db.Debug().Unscoped().Delete(&comment)
}
