package actions

import (
	"BetterDash/configs"
	"BetterDash/models"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(user *models.UserData) string {
	claims := models.Claims{
		Email:    user.Email,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix(),
			Issuer:    "Groszek",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(configs.JWT_SECRET))
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func ParseJWT(r *http.Request) (claims *models.Claims, err error, status int) {

	claims = &models.Claims{}
	cookie, err_c := r.Cookie("JWT")
	if err_c != nil {
		return claims, err, http.StatusBadRequest
	}
	tokenString := cookie.Value

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return configs.JWT_SECRET, nil
	})

	if token != nil {
		return claims, nil, http.StatusOK
	}
	return claims, err, http.StatusMultiStatus
}
