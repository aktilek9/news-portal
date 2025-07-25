package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type JWTService interface {
	GenerateToken(userID int, role string) (string, error)
	ParseToken(token string) (int, string, error)
}

type jwtService struct {
	secretKey []byte
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{secretKey: []byte(secretKey)}
}

func (j *jwtService) GenerateToken(userID int, role string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &CustomClaims{
		UserID: userID,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "news-portal",
		},
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *jwtService) ParseToken(tokenString string) (int, string, error) {
	token, err  := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return j.secretKey, nil
	})
	
	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserID, claims.Role, nil
	}

	return 0, "", jwt.ErrInvalidKey
}
