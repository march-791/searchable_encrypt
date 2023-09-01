package account

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gotest/db"
	"gotest/tool"
	"strings"
	"time"
)

type Claims struct {
	UserID   string
	PassWord string
	jwt.StandardClaims
}
type Auth struct {
	UserID string
	Auth   bool
}

var TokenExpireSeconds int = 7200

func GenToken(userName, passWord string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		userName,
		passWord,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(TokenExpireSeconds)).Unix(),
		},
	})
	return token.SignedString([]byte(tool.DefaultTokenSecretKey))
}

// ParseToken JWT parse token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(tool.DefaultTokenSecretKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token is not available")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token has expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("invalid token")
			} else {
				return nil, fmt.Errorf("token is not available")
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
func Register(info db.UserInfo) (flag int, err error) {
	if info.PassWord == "" || info.UserId == "" {
		return 0, fmt.Errorf("pass or userid cant be empty")
	}
	err = db.InsertUser(info)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return 1, nil
}

func Login(UserId string, UserPassWord string) (string, error) {
	info := db.SearchUserByID(UserId)
	flag := strings.Compare(UserPassWord, info.PassWord)
	if flag != 0 {
		return "", fmt.Errorf("ID or password wrong")
	}
	var err error
	var token string
	token, err = GenToken(UserId, UserPassWord)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return token, nil
}
func Authority(tokenString string) (Auth, error) {

	var auth Auth = Auth{
		UserID: "0",
		Auth:   false,
	}
	claim, err := ParseToken(tokenString)
	if err != nil {
		fmt.Println(err.Error())
		return auth, err
	}

	UserId := claim.UserID
	UserPassWord := claim.PassWord
	info := db.SearchUserByID(UserId)
	flag := strings.Compare(UserPassWord, info.PassWord)
	if flag != 0 {
		return auth, fmt.Errorf("invalid token")
	}
	auth.UserID = UserId
	auth.Auth = true
	return auth, nil
}
