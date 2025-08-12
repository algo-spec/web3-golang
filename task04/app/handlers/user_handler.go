package handlers

import (
	"log"
	"net/http"
	"task04/domain/database"
	"task04/domain/models"
	"task04/pkg/config"
	"task04/pkg/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 用户注册
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[ERROR] Failed to hash password %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(bcryptPassword)
	if err := database.GetDB().Create(&user).Error; err != nil {
		log.Printf("[ERROR] Failed to create user %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// 用户登录，jwt验证
func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbUser models.User
	if err := database.GetDB().Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	config.LoadConfig()
	//生成jwt
	token, err := util.GenerateToken(dbUser.ID, dbUser.Username, config.AppConfig.JwtSecret)
	if err != nil {
		log.Printf("[ERROR] Failed to generate token %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "token": token})
}
