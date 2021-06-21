package lists


type UserCreateList struct {
	Website string `json:"website"`
	Password string `json:"password"`
}

type UserUpdateList struct {
	Website string `json:"website" binding:"required, email"`
	Password string `json:"password"`
}