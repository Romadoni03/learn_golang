package auth_test

import (
	"ecommerce-cloning-app/internal/auth"
	"fmt"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	// expirationTime := time.Now().Local().Add(time.Minute * 30)
	_, err := auth.GenerateJWT("083156490686")
	_, err2 := auth.GenerateJWT("083156490687")
	if err != nil || err2 != nil {
		fmt.Println(err)
	}
	// fmt.Println("token :", token, "expiredAt :", time.Now().Local().Add(time.Minute*30))
	// fmt.Println("token 2 :", token2)
	// fmt.Println(expirationTime)
}

func TestValidateJWT(t *testing.T) {
	token, err := auth.GenerateJWT("083156490686")
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 5)

	claims, errvalidate := auth.ValidateJWT(token)
	if errvalidate != nil {
		fmt.Println(errvalidate)
	}

	fmt.Println(claims)
}

func TestRefreshToken(t *testing.T) {
	for i := 0; i < 10; i++ {
		refresh := auth.GenerateRefreshToken()
		fmt.Println(refresh)
	}

}
