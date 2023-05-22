package auth

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type CustomAuth struct {
	Id       int64
	Username string
	jwtV4.RegisteredClaims
}

var (
	ErrParseJWTTokenFail  = errors.New("获取凭证信息失败")
	ErrParseJWTTokenEmpty = errors.New("凭证信息为空")
)

func GenerateJWTToken(loginInfo CustomAuth, secretKey []byte) string {
	claims := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256, loginInfo)

	signedToken, err := claims.SignedString(secretKey)
	if err != nil {
		return ""
	}

	return signedToken
}

func GetLoginIdByContext(ctx context.Context) (int64, error) {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return 0, ErrParseJWTTokenFail
	}

	mapCliaims := claims.(*jwtV4.MapClaims)
	Id := (*mapCliaims)["Id"].(float64)
	if Id == 0 {
		return 0, ErrParseJWTTokenEmpty
	}
	return int64(Id), nil
}

// Encrypt 密码加密
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare 密码对比
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
