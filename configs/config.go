package configs

import (
	"net/http"
	"os"
)

var FS http.Handler = http.FileServer(http.Dir("./public"))
var JWT_SECRET string = os.Getenv("JWT_SECRET")
