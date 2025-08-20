package routes

func InitRoutes(r *gin.RouterGroup) {
	r.Get("/getUsersById/:userId")
	r.Get("/getUsersByEmail/:userEmail")
	r.Get("/getAllUsers")
	r.Post("/createUser")
	r.Put("/updateUser/:userId")
	r.Delete("/deleteUser/:userId")

}