package auth

type Session struct {
	User  User
	Token Token
}

func CreateSession() (Session, error) {

}

func DeleteSession() error {
}
