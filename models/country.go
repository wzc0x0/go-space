package models

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Population int `json:"population"`
}
