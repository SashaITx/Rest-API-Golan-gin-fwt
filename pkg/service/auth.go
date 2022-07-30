package service

import (
	rest "Rest_API_Golan-gin-fwt"
	"Rest_API_Golan-gin-fwt/pkg/repository"
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

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user rest.User) (int, error) {
	user.Password = s.generationPasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generationPasswordHash(password))
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

func (s *AuthService) generationPasswordHash(password string) string {

	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
