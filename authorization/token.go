package authorization

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sygns13/golang-api/model"
)

//GenerateToken - Generate a JWT token
func GenerateToken(data *model.Login) (string, error) {

	claims := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "go-api-bcs",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

//ValidateToken - Validate a JWT token
func ValidateToken(tokenString string) (model.Claim, error) {

	token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, verifyFunction)

	if err != nil {
		return model.Claim{}, err
	}

	if !token.Valid {
		return model.Claim{}, errors.New("Token is not valid")
	}

	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("No se pudo obtener el claim")
	}

	return *claim, nil

	/* 	if claims, ok := token.Claims.(*model.Claim); ok && token.Valid {
	   		return true, nil
	   	}

	return false, nil
	*/
}

func verifyFunction(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
