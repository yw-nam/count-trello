package api_client

import "github.com/yw-nam/count-trello/models"

type ApiClient interface {
	GetList() []models.List
	GetCards(listId string) []models.Card
}
