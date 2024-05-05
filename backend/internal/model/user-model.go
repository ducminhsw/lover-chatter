package model

type User struct {
	Email        string `json:"email"`
	HashPassword string `json:"hashpassword"`
	Username     string `json:"username"`
	HeartKey     string `json:"heartkey"`
	LoverName    string `json:"lovername"`
	MessagesId   string `json:"messageid"`
}
