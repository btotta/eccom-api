package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/patrickmn/go-cache"
)

var (
	jwtKey           = []byte(os.Getenv("JWT_SECRET"))
	jwtExp, _        = strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
	jwtRefresh       = os.Getenv("JWT_REFRESH_SECRET")
	jwtRefreshExp, _ = strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRES_IN"))
	invalidTokens    = cache.New(12*time.Hour, 10*time.Minute)
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJwtToken(email string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(jwtExp) * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(email string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(jwtRefreshExp) * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtRefresh))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {

	if _, found := invalidTokens.Get(tokenString); found {
		return nil, jwt.ErrSignatureInvalid
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

func ValidateRefreshToken(tokenString string) (*Claims, error) {

	if _, found := invalidTokens.Get(tokenString); found {
		return nil, jwt.ErrSignatureInvalid
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtRefresh), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

func RefreshToken(tokenString string) (string, error) {
	claims, err := ValidateRefreshToken(tokenString)
	if err != nil {
		return "", err
	}

	return GenerateJwtToken(claims.Email)
}

func InvalidateToken(token string) {
	invalidTokens.Set(token, true, cache.DefaultExpiration)
}
