package models

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Password  string `json:"user_password,omitempty"`
	Role      string `json:"user_role"`
}
