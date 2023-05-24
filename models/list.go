package models

type List struct {
	Order  int
	Id     string `json:"id"`
	Name   string `json:"name"`
	Closed bool   `json:"closed"`
}
