package jwt

import (
	"net/http"

	errors "github.com/FernandoCagale/serverless-infra/error"
	"github.com/FernandoCagale/serverless-infra/render"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwtgo "github.com/dgrijalva/jwt-go"
)

//GetConfigJWT options jwt
func GetConfigJWT(secret string) *jwtmiddleware.Options {
	return &jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwtgo.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			render.ResponseError(w, errors.AddUnauthorizedError("Missing Authentication"))
		},
		SigningMethod: jwtgo.SigningMethodHS256,
	}
}
