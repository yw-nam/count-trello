package models

type List struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Closed bool   `json:"closed"`
}
