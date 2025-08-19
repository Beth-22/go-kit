package services

//create jwt tokens on login
//verify tokens for middleware
import (
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)


var jwtKey = []byte("secret-key")

type Claims struct{
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

//create a new token for the user

func GenerateJWT(username, role string) (string, error){
	expirationTime:= time.Now().Add(24*time.Hour)


	claims:= &Claims{
		Username : username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}



//verify and validate the token

func VerifyJWT(tokenStr string) (*Claims,error){
	claims := &Claims{}

	token,err :=jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error){
		return jwtKey, nil

})
  if err !=nil || !token.Valid{
	return nil, errors.New("invalid token")
  }
  return claims, nil
}