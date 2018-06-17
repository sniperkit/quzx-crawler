package rest_api

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"

	"github.com/sniperkit/quzx-crawler/pkg/quzxutil"
	"github.com/sniperkit/quzx-crawler/pkg/rest-api/controllers"
	"github.com/sniperkit/quzx-crawler/pkg/rest-api/routers"
)

func Serve() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	router := mux.NewRouter()
	router = routers.SetAuthRoute(router)

	apiRoutes := routers.InitRoutes()

	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(controllers.ValidateTokenMiddleware),
		negroni.Wrap(apiRoutes),
	))

	server := negroni.Classic()
	server.Use(c)
	server.UseHandler(router)
	server.Run("0.0.0.0:" + quzxutil.GetParameter("PORT"))
}
