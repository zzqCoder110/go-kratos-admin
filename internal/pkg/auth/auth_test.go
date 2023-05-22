package auth

import (
	"fmt"
	"testing"
)

const JwtKey = "12342314"

func TestGenerateToken(t *testing.T) {
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpdHlJZCI6IjEifQ.4gary3Ffo6dVnqtEtNJLAqcJbAqaU_J2zVbkWUWBN7w"
	//authn := &Auth{}
	//if err := authn.ParseJWTToken(token, []byte("gosimgosimgogogosimsimsim")); err != nil {
	//	fmt.Println("解析失败")
	//}
	//fmt.Printf("data %v \n", authn.Subject)
	custom := CustomAuth{
		Id:       999,
		Username: "custom",
	}
	token := GenerateJWTToken(custom, []byte(JwtKey))
	fmt.Println(token)
}
