package auth

type User struct {
	Email    string `json: "email"`
	Password string `json: "password, omitempty"`
	Number   int    `json: "user_number, omitempty"`
}
