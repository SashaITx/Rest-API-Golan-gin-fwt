package service

import (
	rest "Rest_API_Golan-gin-fwt"
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

type Auth interface {
	CreateUser(user rest.User) (int, error)
	GetUser(username, password string) (rest.User, error)
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

func (s *AuthService) CreateUser(user rest.User) (int, error) {
	user.Password = s.generationPasswordHash(user.Password)

	return s.auth.CreateUser(user)
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
