package dtos

type userResponse struct {
	Id 	 string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age	  int8    `json:"age"`
}