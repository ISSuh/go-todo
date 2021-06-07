package db

type User struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type Account struct {
	User       User `json: "user"`
	AccountNum int  `json: "account_num"`
}
