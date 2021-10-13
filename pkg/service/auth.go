package service

import (
	"crawl-data/conf"
	"crawl-data/pkg/model"
	"crawl-data/pkg/repo"
	"crawl-data/pkg/utils"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTService struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

func NewJWTService() IJWTService {
	return &JWTService{}
}

type IJWTService interface {
	GenJWTToken(userID uuid.UUID) (string, error)
}

type AuthService struct {
	Repo repo.IRepo
}

func NewAuthService(repo repo.IRepo) IAuthService {
	return &AuthService{
		Repo: repo,
	}
}

type IAuthService interface {
	SignUp(email string, password string) (*model.User, error)
	CheckUserPassword(email string, password string) (*model.User, error)
}

func (s *AuthService) SignUp(email string, password string) (*model.User, error) {
	//Logic base on diagram
	user, err := s.Repo.CheckEmail(email)
	//err != nil, not ErrNotFound
	if err != nil && !utils.IsErrNotFound(err) {
		//Case: db error
		return nil, err
	}
	if user != nil {
		//user existed
		//return err
		return nil, fmt.Errorf("email existed")
	}

	//create user

	hashpassword := password + utils.SystemHashKey
	newUser := &model.User{
		Email: email,
		//Ma hoa pass trc
		Password: hashpassword,
	}

	user, err = s.Repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) CheckUserPassword(email string, password string) (*model.User, error) {
	//get user by email
	user, err := s.Repo.CheckEmail(email)
	if err != nil {
		return nil, err
	}

	//check request password == user.password
	//Ma hoa pass
	hashpassword := password + utils.SystemHashKey
	if hashpassword != user.Password {
		return nil, fmt.Errorf("wrong password")
	}
	//pass correct
	//return JWT token
	return user, nil
}

func (s *JWTService) GenJWTToken(userID uuid.UUID) (string, error) {
	claims := &JWTService{
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
