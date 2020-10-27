package actions

import (
	"BetterDash/data"
	"BetterDash/models"
	"net/http"

	"github.com/dchest/uniuri"
)

type UserR models.UserData

func New() *UserR {
	return new(UserR)
}

func (user *UserR) Register() (int, string) {
	valid, response_string := user.ValidateUser()
	if valid {
		_user := models.UserData(*user)
		user.ID = data.SetID()
		if !ValidEmail(&_user) {
			return http.StatusBadRequest, "invalid email"
		}

		user.Verify = uniuri.New()
		user.Active = false
		_user = models.UserData(*user)
		create := data.CreateUser(&_user)
		if create {
			return http.StatusOK, response_string
		} else {
			return http.StatusBadRequest, "db error ID validation"
		}

	} else {
		return http.StatusBadRequest, response_string
	}
}

func ValidEmail(user *models.UserData) bool {
	valid_email := data.IsExist(user)

	if valid_email {
		return false
	}
	return true
}

func (user *UserR) Auth() (int, string) {
	_user := models.UserData(*user)
	is_exist := data.IsExist(&_user)

	if !is_exist {
		return http.StatusForbidden, "invalid email or password"
	}

	if user.Password != data.GetPassword(&_user) {
		return http.StatusForbidden, "invalid email or password"
	}

	if !data.IsActive(user.Email) {
		return http.StatusUnauthorized, "account unactive"
	}

	return http.StatusOK, "success"
}
