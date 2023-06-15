package auth

import (
	"errors"

	"startup-campaign/utils"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToke(userId int) (string, error)
	ValidateToken(encodeToken string) (*jwt.Token, error)
}

type jwtService struct{}

func NewService() *jwtService {
	return &jwtService{}
}

func (j *jwtService) GenerateToke(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecretKeyString, _ := utils.EnvVariabel("JWT_SECRET_KEY")
	jwtSecretKeyByte := []byte(jwtSecretKeyString)

	signedToken, err := token.SignedString(jwtSecretKeyByte)
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

		jwtSecretKeyString, _ := utils.EnvVariabel("JWT_SECRET_KEY")

		return []byte(jwtSecretKeyString), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
