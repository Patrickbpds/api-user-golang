package routes

import (
	"api-user-golang/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUsersById/:userId", controller.FindUserById)
	r.GET("/getUsersByEmail/:userEmail", controller.FindUserByEmail)
	r.GET("/getAllUsers", controller.GetAllUsers)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

}