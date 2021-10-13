package service

import (
	"crawl-data/conf"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type AuthService struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

func NewAuthService() IAuthService {
	return &AuthService{}
}

type IAuthService interface {
	GenJWTToken(userID uuid.UUID) (string, error)
}

func (s *AuthService) GenJWTToken(userID uuid.UUID) (string, error) {
	claims := &AuthService{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "HuyHoang",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(conf.LoadEnv().SecretKey))
	if err != nil {
		return "", err
	}
	fmt.Println(conf.LoadEnv().SecretKey)
	return signedToken, nil
}
