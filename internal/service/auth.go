package service

import (
	"Rest_API_Golan-gin-fwt/internal/repository"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "sfh37s53264jsdfb123jn"
	salt       = "sdflkhs;lsdv;lm123409792245076"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type User struct {
	Name         string
	Username     string
	PasswordHash string
}

type Auth interface {
	CreateUser(user repository.User) (int, error)
	GetUser(username, password string) (repository.User, error)
}

type AuthService struct {
	auth Auth
}

func NewAuthService(auth Auth) *AuthService {
	return &AuthService{auth: auth}
}

func (s *AuthService) generationPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) CreateUser(user User) (int, error) {
	user = s.CreateServiceUserPasswordHash(user.Name, user.Username, user.PasswordHash)
	return s.auth.CreateUser(repository.User{
		Name:         user.Name,
		Username:     user.Username,
		PasswordHash: user.PasswordHash})
}

func (s *AuthService) CreateServiceUserPasswordHash(name, username, password string) User {
	passwordHash := s.generationPasswordHash(password)
	return User{
		Name:         name,
		Username:     username,
		PasswordHash: passwordHash}
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.auth.GetUser(username, s.generationPasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}
