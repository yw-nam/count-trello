package counter

import (
	"time"

	"github.com/yw-nam/count-trello/api"
	"github.com/yw-nam/count-trello/models"
)

type counter struct {
	targetLabel string
	apiClient   api.Client
	baseDate    time.Time
}

func NewCounter(api api.Client, targetLabel string, baseDate time.Time) *counter {
	return &counter{
		apiClient:   api,
		targetLabel: targetLabel,
		baseDate:    baseDate,
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
		Order:    list.Order,
		ListName: list.Name,
		Total:    count,
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

func (a *counter) GetCardCountsByWeeks() models.CardCountSlice {
	results := []models.CardCount{}
	lists := a.apiClient.GetList()
	for _, list := range lists {
		total, byWeek := a.getCardCountByWeeks(list)
		result := models.CardCount{
			Order:    list.Order,
			ListName: list.Name,
			Total:    total,
			ByWeek:   byWeek,
		}
		results = append(results, result)
	}
	return results
}

func (a *counter) getCardCountByWeeks(list models.List) (int, map[int]int) {
	totalCount := 0
	weekCount := map[int]int{}
	cards := a.apiClient.GetCards(list.Id)
	for _, card := range cards {
		if !card.HasLabel(a.targetLabel) {
			continue
		}
		totalCount += 1
		act := a.apiClient.GetCreateAction(card)
		week := weeksAgo(a.baseDate, act.Date)
		weekCount[week] += 1
	}
	return totalCount, weekCount
}

func weeksAgo(baseDate, targetDate time.Time) int {
	duration := time.Since(targetDate)
	weeksAgo := int(duration.Hours() / 24 / 7)
	return weeksAgo
}
