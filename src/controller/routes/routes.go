package routes

import (
	"api-user-golang/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
	) {
	r.GET("/getUsersById/:userId", userController.FindUserById)
	r.GET("/getUsersByEmail/:userEmail", userController.FindUserByEmail)
	r.GET("/getAllUsers", userController.GetAllUsers)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

}