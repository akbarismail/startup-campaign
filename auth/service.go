package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToke(userId int) (string, error)
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
