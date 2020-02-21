package auth

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

// TokenAuth is used fir auth
var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("my_super_secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{"user_id": "1"})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}
