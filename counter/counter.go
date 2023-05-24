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
		list.Order = i
		go a.getCardCount(list, ch)
	}

	results := []models.CardCount{}
	for i := 0; i < len(lists); i++ {
		res := <-ch
		results = append(results, res)
	}
	return results
}

func (a *counter) getCardCount(list models.List, ch chan<- models.CardCount) {
	cards := a.apiClient.GetCards(list.Id)
	count := a.countCardsWithLabel(cards)
	res := models.CardCount{
		Order:     list.Order,
		ListName:  list.Name,
		CardCount: count,
	}
	ch <- res
}

func (a *counter) countCardsWithLabel(cards []models.Card) int {
	count := 0
	for _, c := range cards {
		if c.HasLabel(a.targetLabel) {
			count += 1
		}
	}
	return count
}
