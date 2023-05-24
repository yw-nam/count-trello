package api

import "github.com/yw-nam/count-trello/models"

type Client interface {
	GetList() []models.List
	GetCards(listId string) []models.Card
}
