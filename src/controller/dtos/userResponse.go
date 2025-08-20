package dtos

type UserResponse struct {
	Id 	 string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age	  int8    `json:"age"`
}