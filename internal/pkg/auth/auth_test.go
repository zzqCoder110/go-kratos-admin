package auth

import (
	"fmt"
	"testing"
)

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpdHlJZCI6IjEifQ.4gary3Ffo6dVnqtEtNJLAqcJbAqaU_J2zVbkWUWBN7w"
	authn := &Auth{}
	if err := authn.ParseJWTToken(token, []byte("gosimgosimgogogosimsimsim")); err != nil {
		fmt.Println("解析失败")
	}
	fmt.Printf("data %v \n", authn.Subject)
}
