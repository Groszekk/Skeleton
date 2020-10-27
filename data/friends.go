package data

import "log"

func Search(nick string) [20]string {
	row, err := db.Query(`SELECT nick FROM users WHERE nick like "` + nick + `%"`)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()
	counter := 0
	var nicks [20]string // todo: change this shit!
	for row.Next() {
		row.Scan(&nicks[counter])
		counter++
	}
	return nicks
}
