package infra

import (
	"net/http"

	"github.com/FernandoCagale/serverless-infra/error"
	"github.com/FernandoCagale/serverless-infra/render"
)

//NotFoundHandler NotFound
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.ResponseError(w, error.AddNotFoundError("Not Found"))
	})
}

//NotAllowedHandler NotAllowed
func NotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.ResponseError(w, error.AddMethodNotAllowedError("Method Not Allowed"))
	})
}

//UnauthorizeddHandler Unauthorize
func UnauthorizeddHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.ResponseError(w, error.AddUnauthorizedError("Missing Authentication"))
	})
}
