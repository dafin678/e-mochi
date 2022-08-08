package app

import (
	"e-mochi-app/app/controllers"

	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/register", controllers.RegisterPage).Methods("GET")
	server.Router.HandleFunc("/login", controllers.LoginPage).Methods("GET")
}
