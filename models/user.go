package models

type UserData struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=5,max=42"`
	RePasword string `json:"repassword"`
	Nick      string `json:"nick"`
	ID        uint32
	Verify    string
	Active    bool
	Cookie    string
	Friends   []uint32 // friend's ID's
	Posts     string   // todo: post^struct
}

// type UserLogin struct {
// 	Login    string
// 	Password string
// }
