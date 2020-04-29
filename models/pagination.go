package models

type PaginationDto struct {
	Current int `form:"current,default=1" json:"current"`
	PageSize int `form:"page_size,default=8" json:"page_size"`
}

type Pagination struct {
	PaginationDto
	Total int `json:"total"`
	List interface{} `json:"list"`
}
