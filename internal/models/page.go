package models

type Page struct {
	Content       interface{} `json:"content"`
	TotalElements int         `json:"totalElements"`
	TotalPages    int         `json:"totalPages"`
	Size          int         `json:"size"`
	Number        int         `json:"number"`
}
