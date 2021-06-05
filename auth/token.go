package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	AccessToken  string `json: "access_token`
	RefreshToken string `json: "refresh_token"`
}

const TEST_ACESS_TOKEN_SIGN = "TEST_ACESS_TOKEN_SIGN"
const TEST_REFRESH_TOKEN_SIGN = "TEST_REFRESH_TOKEN_SIGN"

func CreateToken(user *User) error {
	token := Token{
		AccessToken:  "",
		RefreshToken: "",
	}

	var err error
	token.AccessToken, err = createAccessToken(user.Email, user.Numbr)
	if err != nil {
		return err
	}

	token.RefreshToken, err = createRefreshToken(user.Email, user.Numbr)
	if err != nil {
		return err
	}

	return nil
}

func createAccessToken(userEmail string, userNumber int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = userEmail
	claims["user_number"] = userNumber
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedAccessToken, err := accessToken.SignedString([]byte(TEST_ACESS_TOKEN_SIGN))
	if err != nil {
		return "", err
	}
	return signedAccessToken, nil
}

func createRefreshToken(userEmail string, userNumber int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = userEmail
	claims["user_number"] = userNumber
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedRefreshToken, err := refreshToken.SignedString([]byte(TEST_REFRESH_TOKEN_SIGN))
	if err != nil {
		return "", err
	}
	return signedRefreshToken, nil
}
