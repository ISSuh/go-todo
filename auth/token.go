package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const TEST_ACESS_TOKEN_SIGN = "TEST_ACESS_TOKEN_SIGN"
const TEST_REFRESH_TOKEN_SIGN = "TEST_REFRESH_TOKEN_SIGN"

type Claim struct {
	Email string
	jwt.StandardClaims
}

type Token struct {
	AccessToken  string `json: "access_token`
	RefreshToken string `json: "refresh_token"`
}

func CreateToken(email string) (*Token, error) {
	var err error
	accessToken := ""
	refreshToken := ""

	accessToken, err = createAccessToken(email)
	if err != nil {
		return nil, err
	}

	refreshToken, err = createRefreshToken(email)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func ExtractToken(req *http.Request) (string, error) {
	signedToken := extractTokenString(req)
	claims, err := verifyToken(signedToken)
	if err != nil {
		return "", err
	}

	return claims.Email, nil
}

func createAccessToken(userEmail string) (string, error) {
	claims := &Claim{
		Email: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedAccessToken, err := accessToken.SignedString([]byte(TEST_ACESS_TOKEN_SIGN))
	if err != nil {
		return "", err
	}
	return signedAccessToken, nil
}

func createRefreshToken(userEmail string) (string, error) {
	claims := &Claim{
		Email: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(TEST_REFRESH_TOKEN_SIGN))
	if err != nil {
		return "", err
	}
	return signedRefreshToken, nil
}

func verifyToken(signedToken string) (*Claim, error) {
	claims, err := parseSignedToken(signedToken)
	if err != nil {
		return nil, err
	}

	if err = checkTokenExpires(claims); err != nil {
		return nil, err
	}

	return claims, nil
}

func extractTokenString(r *http.Request) string {
	tockenHeader := r.Header.Get("Authorization")

	strArr := strings.Split(tockenHeader, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func parseSignedToken(signedToken string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(TEST_ACESS_TOKEN_SIGN), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if claims, res := token.Claims.(*Claim); res {
		return claims, nil
	}

	return nil, errors.New("Couldn't parse claims")
}

func checkTokenExpires(claims *Claim) error {
	if claims.ExpiresAt < time.Now().Unix() {
		return errors.New("Token is expired")
	}
	return nil
}
