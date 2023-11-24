package entities

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
