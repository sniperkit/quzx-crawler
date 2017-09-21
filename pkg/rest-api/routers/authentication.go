package routers

import (
	"github.com/gorilla/mux"
	"github.com/demas/cowl-go/pkg/rest-api/controllers"
)

func SetAuthRoute(router *mux.Router) *mux.Router {

	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")
	return router
}
