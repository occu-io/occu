package router

import (
	"auth"
	"service"
	"settings"

	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc(settings.APIPath+"/token/auth", auth.Authenticate).Methods("POST")
	router.HandleFunc(settings.APIPath+"/token/refresh", auth.ValidateMiddleware(auth.RefreshToken)).Methods("GET")

	router.HandleFunc(settings.APIPath+"/test", auth.ValidateMiddleware(service.TestEndpoint)).Methods("GET")
	return router
}
