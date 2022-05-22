package utils

import (
	"crypto/rsa"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"admin/models"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

func init() {
	var err error

	// Carga de llaves publica y privada
	signKey, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}
	verifyKey, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(signKey)
	if err != nil {
		log.Fatal("Error parsing private key")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKey)
	if err != nil {
		log.Fatal("Error parsing public key")
	}
}

func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			Issuer:    "Issue Standard Claim",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("Error signing Token")
	}
	return result
}

func ValidateToken(r *http.Request) error {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return checkTokenError(err)
	}

	if token.Valid {

		// user := token.Claims.(*models.Claim).User
		return nil
	} else {
		// w.WriteHeader(http.StatusUnauthorized)
		// fmt.Println("Invalid Token")
		return models.ErrorValidationToken
	}
}

func GetUserFromToken(r *http.Request) (models.User, error) {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	var user models.User
	if err != nil {
		return user, checkTokenError(err)
	}

	if token.Valid {
		user = token.Claims.(*models.Claim).User
		return user, nil
	} else {
		return user, models.ErrorValidationToken
	}
}

func checkTokenError(err error) error {
	switch err.(type) {

	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return models.ErrorExpiredToken
		case jwt.ValidationErrorSignatureInvalid:
			return models.ErrorSign
		default:
			return models.ErrorValidationToken
		}

	default:
		return models.ErrorValidationToken
	}
}
