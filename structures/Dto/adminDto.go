package dto



type CreateUserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ChangePassword struct{
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

