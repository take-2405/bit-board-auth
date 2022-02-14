package middleware

import (
	"github.com/dgrijalva/jwt-go"
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
	//TODO ここでローカルの秘密鍵を使用する
	tokenString, _ := token.SignedString([]byte("SECRET_KEY"))
	return tokenString
}
