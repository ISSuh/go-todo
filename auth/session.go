package auth

import "errors"

type Session struct {
	User  User
	Token Token
}

func CreateSession(user User) (*Session, error) {
	token, err := CreateToken(&user)
	if err != nil {
		return nil, errors.New("Fail creating token")
	}

	return &Session{
		User:  user,
		Token: *token,
	}, nil
}
