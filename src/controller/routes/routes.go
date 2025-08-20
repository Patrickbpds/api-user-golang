package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUsersById/:userId")
	r.GET("/getUsersByEmail/:userEmail")
	r.GET("/getAllUsers")
	r.POST("/createUser")
	r.PUT("/updateUser/:userId")
	r.DELETE("/deleteUser/:userId")

}