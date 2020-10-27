package actions

import (
	"BetterDash/data"
)

func Search(nick string) [20]string {
	nicks := data.Search(nick)
	// fmt.Println(nicks)
	return nicks
}
