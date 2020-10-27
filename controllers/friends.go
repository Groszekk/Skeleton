package controllers

import (
	"BetterDash/actions"
	"fmt"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()["nick"]
	nick := string(param[0])
	nicks := actions.Search(nick)
	for i := 0; i < len(nicks); i++ {
		fmt.Fprintln(w, nicks[i])
	}
}
