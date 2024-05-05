package model

type Request struct {
	Id     string `json:"id"`
	From   string `json:"from"`
	Target string `json:"target"`
	Note   string `json:"note"`
}
