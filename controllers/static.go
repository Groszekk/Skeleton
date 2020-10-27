package controllers

import (
	"BetterDash/configs"
	"net/http"
)

func StaticRegister(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "/html/register.html"
	configs.FS.ServeHTTP(w, r)
}

func StaticLogin(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "/html/login.html"
	configs.FS.ServeHTTP(w, r)
}

func StaticHome(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "/html/home.html"
	configs.FS.ServeHTTP(w, r)
}
