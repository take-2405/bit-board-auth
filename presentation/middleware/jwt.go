package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateJwt(userID string) string {
	// Claimsオブジェクトの作成
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//fmt.Printf("Header: %#v\n", token.Header) // Header: map[string]interface {}{"alg":"HS256", "typ":"JWT"}
	//fmt.Printf("Claims: %#v\n", token.Claims) // CClaims: jwt.MapClaims{"exp":1634051243, "user_id":12345678}

	// トークンに署名を付与
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return tokenString
}
