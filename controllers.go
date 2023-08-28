package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser.Password = hashedPassword
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	if err := uc.db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("userId")
	userIDInt := parseInt(userID)

	user, err := uc.getUserByID(userIDInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userID := c.Param("userId")
	userIDInt := parseInt(userID)

	user, err := uc.getUserByID(userIDInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	user.UpdatedAt = time.Now()

	if err := uc.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("userId")
	userIDInt := parseInt(userID)

	user, err := uc.getUserByID(userIDInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := uc.db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (uc *UserController) getUserByID(id int) (*User, error) {
	var user User
	if err := uc.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (controller *UserController) LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := controller.db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
// ... other user-related functions ...

type PhotoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{db: db}
}

func (pc *PhotoController) CreatePhoto(c *gin.Context) {
	// Implement create photo logic here...
}

func (pc *PhotoController) GetPhotos(c *gin.Context) {
	// Implement get photos logic here...
}

func (pc *PhotoController) GetPhotoByID(c *gin.Context) {
	// Implement get photo by ID logic here...
}

func (pc *PhotoController) UpdatePhoto(c *gin.Context) {
	// Implement update photo logic here...
}

func (pc *PhotoController) DeletePhoto(c *gin.Context) {
	// Implement delete photo logic here...
}

// ... your other utility functions ...

func parseInt(s string) int {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		return 0
	}
	return n
}
