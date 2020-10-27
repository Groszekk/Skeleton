package controllers

import (
	"BetterDash/actions"
	"BetterDash/data"
	"BetterDash/models"
	"BetterDash/smtp"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := actions.New()
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid json")
		return
	}

	status, response := user.Register()
	w.WriteHeader(status)
	fmt.Fprint(w, response)
	if status == http.StatusOK {
		_user := models.UserData(*user)
		go smtp.Verify(&_user)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := actions.New()
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid json")
		return
	}

	status, _ := user.Auth()
	switch status {
	case http.StatusOK:
		_user := models.UserData(*user)
		jwt := actions.CreateJWT(&_user)
		jwt_cookie := http.Cookie{Name: "JWT", Value: jwt, Path: "/dash", HttpOnly: true}
		http.SetCookie(w, &jwt_cookie)
		w.WriteHeader(status)
		fmt.Fprint(w)
		return

	case http.StatusUnauthorized:
		w.WriteHeader(status)
		fmt.Fprint(w)
		return
	}
	w.WriteHeader(status)
	fmt.Fprint(w, "something wrong")
}

func Verify(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()["token"]
	token := string(param[0])
	verify := data.GetVerifyStrings()
	if token == string(verify) {
		fmt.Fprint(w, "Your account was activated !!!")
		data.Activate(string(verify))
		return
	}
	fmt.Fprint(w, "Something wrong with your token")
}

func DashBoard(w http.ResponseWriter, r *http.Request) {
	parsed_jwt, err, status := actions.ParseJWT(r)
	if err != nil {
		log.Fatal("\n", err)
	}

	if status != http.StatusOK {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		fmt.Fprint(w)
		return
	}

	w.WriteHeader(status)
	tmpl := template.Must(template.ParseFiles("public/html/dashboard/index.html"))
	tmpl.Execute(w, *parsed_jwt)
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
}
