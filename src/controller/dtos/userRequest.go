package dtos

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=100"`
	Password string `json:"password" binding:"required,min=8,containsany=@#$%&*"`
	Email    string `json:"email" binding:"required,email"`
	Age      int8   `json:"age" binding:"required,min=0,max=120"`
}
