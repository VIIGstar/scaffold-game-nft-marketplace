package entities

type User struct {
	Investor
	BaseGameEntity
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
