package handlers

import (
	"log"
	"net/http"
	"task04/domain/database"
	"task04/domain/models"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
}

// 创建评论
func CreateComment(c *gin.Context) {
	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "请先登录"})
		return
	}

	comment := models.Comment{
		PostID:  req.PostID,
		Content: req.Content,
		UserID:  userID.(uint),
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		log.Printf("[ERROR] create comment failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "评论创建失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 获取评论
func GetComment(c *gin.Context) {
	postID := c.Param("post_id")
	var comments []models.Comment
	if err := database.DB.Where("post_id = ?", postID).Preload("User").Find(&comments).Error; err != nil {
		log.Printf("[ERROR] get comments failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
		return
	}

	var res []CommentResponse
	for _, comment := range comments {
		res = append(res, CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			PostID:    comment.PostID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": res})
}
