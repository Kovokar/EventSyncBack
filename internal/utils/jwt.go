package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtUtil struct {
	secretKey string
	issure    string
}

func NewJtw() *jwtUtil {
	return &jwtUtil{
		secretKey: "Pepe moreno tenho fome! Cante meu fi cante",
		issure:    "volei-api",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (u *jwtUtil) GenenateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    u.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(u.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtUtil) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid Token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}
