package models

import "time"

const (
	ActionCreateCard string = "createCard"
)

type Action struct {
	Id   string    `json:"id"`
	Type string    `json:"type"`
	Date time.Time `json:"date"`
}
