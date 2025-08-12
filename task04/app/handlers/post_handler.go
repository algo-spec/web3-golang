package handlers

import (
	"log"
	"net/http"
	"task04/domain/database"
	"task04/domain/models"

	"github.com/gin-gonic/gin"
)

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 创建文章
func CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数异常"})
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "请先登录"})
		return
	}

	user := models.User{}
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  user.ID,
	}
	if err := database.DB.Create(&post).Error; err != nil {
		log.Printf("[ERROR] create post failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

// 获取文章列表
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Preload("User").Find(&posts).Error; err != nil {
		log.Printf("[ERROR] get posts failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": posts})
}

// 获取文章
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.Preload("User").First(&post, id).Error; err != nil {
		log.Printf("[ERROR] get post failed: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": post})
}

// 更新文章
func UpdatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数异常"})
		return
	}
	isValid, post := valid(c)

	if !isValid {
		return
	}
	post.Title = req.Title
	post.Content = req.Content
	if err := database.DB.Save(&post).Error; err != nil {
		log.Printf("[ERROR] update post failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// 删除文章
func DeletePost(c *gin.Context) {
	isValid, post := valid(c)
	if !isValid {
		return
	}

	if err := database.DB.Delete(&post).Error; err != nil {
		log.Printf("[ERROR] delete post failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func valid(c *gin.Context) (bool, models.Post) {
	var post models.Post
	id := c.Param("id")
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "请先登录"})
		return false, post
	}

	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return false, post
	}

	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权操作"})
		return false, post
	}
	return true, post
}
