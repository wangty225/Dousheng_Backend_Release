package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	userJwt := NewJWT()

	token, err := userJwt.CreateToken(CustomClaims{
		1460162896,
		jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(time.Second * 5).Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 7500).Unix(),
			Issuer:    "dousheng",
		},
	})

	token2, err := userJwt.CreateToken(CustomClaims{
		2160924050,
		jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(time.Second * 5).Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 7500).Unix(),
			Issuer:    "dousheng",
		},
	})

	fmt.Printf("%v\n", token)
	fmt.Printf("%v\n", token2)

	if err != nil {
		t.Fatalf("create token error %v", err)
	}

	_, err = userJwt.ParseToken(token)

	if err != nil {
		t.Fatalf("token verified error %v", err)
	}

	otherJwt := NewJWT()
	c, err := otherJwt.ParseToken(token)

	if err != nil {
		t.Fatalf("token verified error %v", err)
	}
	//expirationTime := time.Unix(int64(c["expat"].(float64)), 0)
	fmt.Printf("exp at : %v\n", c)

	time.Sleep(time.Second * 7)

	_, err = userJwt.ParseToken(token)

	//if err == nil {
	//	t.Fatalf("token expired but not got error")
	//}

}

func TestJWT_GetIdFromToken(t *testing.T) {
	// Ipc : 1460162896 : wangty002
	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MTQ2MDE2Mjg5NiwiZXhwIjoxNzIwMjE3OTQ1LCJpc3MiOiJkb3VzaGVuZyJ9.gb1zbmucGAE2fOhTieJHEBwAhwCTv_QQXVbvbAiWIpc"

	// 3oU : 2160924050 : wangty006
	//tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MjE2MDkyNDA1MCwiZXhwIjoxNzIwNDU4MzMyLCJpc3MiOiJkb3VzaGVuZyJ9.EzIcDAxciGlzBELiR70tqzFnW9wNTLat24vAxarO3oU"

	testJwt := NewJWT()
	userID, err := testJwt.GetIdFromToken(tk)
	if err != nil {
		message := "token解析错误，无法获取用户ID"
		fmt.Printf("%v\n", message)
	}
	fmt.Printf("%v\n", userID)
}
