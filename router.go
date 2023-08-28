package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *UserController, photoController *PhotoController) *gin.Engine {
	r := gin.Default()

	authMiddleware := authMiddleware()

	r.POST("/users/register", userController.RegisterUser)
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
