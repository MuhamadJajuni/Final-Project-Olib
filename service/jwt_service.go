package service

import (
	"final-project-olib/config"
	"final-project-olib/model"
	"final-project-olib/model/dto"
	"final-project-olib/utils"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	CreateToken(author model.Borrower) (dto.AuthResponDto, error)
	ValidateToken(token string) (jwt.MapClaims, error)
}

type jwtService struct {
	co config.TokenConfig
}

// Create Token
func (j *jwtService) CreateToken(author model.Borrower) (dto.AuthResponDto, error) {
	claims := utils.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.co.IssuerName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.co.ExpiresTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Name:  author.Name,
		Email: author.Email,
	}

	token := jwt.NewWithClaims(j.co.SigningMethod, claims)
	ss, err := token.SignedString(j.co.SignatureKey)
	if err != nil {
		return dto.AuthResponDto{}, fmt.Errorf("failed create access token")
	}
	return dto.AuthResponDto{Token: ss}, nil
}

// Verify Token
func (j *jwtService) ValidateToken(tokenHeader string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		return j.co.SignatureKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to verify token when parsing")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to verify token when claims")
	}
	return claims, nil
}

func NewJwtService(c config.TokenConfig) JwtService {
	return &jwtService{co: c}
}
