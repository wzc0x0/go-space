package models

type SpMovie struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Year int `json:"year"`
	Desc string `json:"desc"`
	Country string `json:"country"`
	Star float64 `json:"star"`
	Comment int `json:"comment"`
	Quote string `json:"quote"`
	ImgUrl string `json:"img_url"`
}