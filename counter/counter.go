package counter

import (
	"github.com/yw-nam/count-trello/api_client"
	"github.com/yw-nam/count-trello/models"
)

type counter struct {
	targetLabel string
	apiClient   api_client.ApiClient
}

func NewCounter(api api_client.ApiClient, targetLabel string) *counter {
	return &counter{
		apiClient:   api,
		targetLabel: targetLabel,
	}
}

func (a *counter) GetCardCounts() models.CardCountSlice {
	lists := a.apiClient.GetList()

	ch := make(chan models.CardCount, len(lists))
	for i, list := range lists {
		go a.getCardCount(i, list, ch)
	}

	results := []models.CardCount{}
	for i := 0; i < len(lists); i++ {
		res := <-ch
		results = append(results, res)
	}
	return results
}

func (a *counter) getCardCount(order int, list models.List, ch chan<- models.CardCount) {
	cards := a.apiClient.GetCards(list.Id)
	count := models.CountCardsHavingLabel(a.targetLabel, cards)
	res := models.CardCount{
		Order:     order,
		ListName:  list.Name,
		CardCount: count,
	}
	ch <- res
}
