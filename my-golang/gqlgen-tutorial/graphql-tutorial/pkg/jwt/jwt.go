package jwt

import (
	"log"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

// 秘密鍵
var (
	SecretKey = []byte("secret")
)

// GenerateTokenはjwtトークンを生成する
// そのclaimにユーザ名を割り当てて、返す
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	/* クレームを保存するためのマップを作成する */
	claims := token.Claims.(jwt.MapClaims)
	/* クレーム・トークンを設定する */
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

// ParseTokenはjwtトークンを解析し、ユーザ名を返す
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
