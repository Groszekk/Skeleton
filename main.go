package main

import (
	"BetterDash/controllers"
	"BetterDash/scripts"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	scripts.InitDB()
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.PathPrefix("/home").Handler(http.HandlerFunc(controllers.StaticHome)).Methods("GET")
	r.HandleFunc("/user/register", controllers.Register).Methods("POST")
	r.HandleFunc("/user/login", controllers.Login).Methods("POST")
	r.HandleFunc("/verify", controllers.Verify).Methods("GET")
	r.HandleFunc("/dash", controllers.DashBoard).Methods("GET")
	r.PathPrefix("/register").Handler(http.HandlerFunc(controllers.StaticRegister)).Methods("GET")
	r.PathPrefix("/login").Handler(http.HandlerFunc(controllers.StaticLogin)).Methods("GET")
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	r.HandleFunc("/friends/search", controllers.Search).Methods("GET")
	http.ListenAndServe(":8080", handlers.CORS()(r))
}
