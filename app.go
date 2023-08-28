package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = setupDatabase()
	userController := NewUserController(db)
	photoController := NewPhotoController(db)

	r := setupRouter(userController, photoController)
	r.Run(":8080")
}


func setupRouter(userController *UserController, photoController *PhotoController) *gin.Engine {
	r := gin.Default()

	authMiddleware := authMiddleware()

	r.POST("/users/register", userController.RegisterUser)
	r.POST("/users/login", userController.LoginUser)
	r.GET("/users/:userId", userController.GetUserByID)
	r.PUT("/users/:userId", userController.UpdateUser)
	r.DELETE("/users/:userId", userController.DeleteUser)

	r.POST("/photos", authMiddleware, photoController.CreatePhoto)
	r.GET("/photos", authMiddleware, photoController.GetPhotos)
	r.GET("/photos/:photoId", authMiddleware, photoController.GetPhotoByID)
	r.PUT("/photos/:photoId", authMiddleware, photoController.UpdatePhoto)
	r.DELETE("/photos/:photoId", authMiddleware, photoController.DeletePhoto)

	return r
}
