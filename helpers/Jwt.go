package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("secret") // kunci rahasia

func GenerateToken(UserID uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = UserID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(SecretKey)
}

func VerifyToken(tokenString string)(jwt.MapClaims, error){
	token, err := jwt.Parse(tokenString, func(token * jwt.Token)(interface {},error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signing method")
		}
		return SecretKey,nil
	})

	if err != nil {return nil, err}
	if !token.Valid {return nil, fmt.Errorf("invalid token")}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {return nil, fmt.Errorf("invalid claims")}
	return claims,nil
}
