package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToke(userId int) (string, error)
	ValidateToken(encodeToken string) (*jwt.Token, error)
}

type jwtService struct{}

var jwtSecretKey = []byte("startup_campaign_S3CR3T")

func NewService() *jwtService {
	return &jwtService{}
}

func (j *jwtService) GenerateToke(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (j *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodeToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
