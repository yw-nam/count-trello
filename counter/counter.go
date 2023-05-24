package counter

import (
	"fmt"
	"log"
	"time"

	"github.com/yw-nam/count-trello/api"
	"github.com/yw-nam/count-trello/models"
)

type counter struct {
	targetLabel string
	apiClient   api.Client
}

func NewCounter(api api.Client, targetLabel string) *counter {
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

func (a *counter) GetCardCountsByWeeks() {
	lists := a.apiClient.GetList()

	for i, list := range lists {
		list.Order = i
		fmt.Printf("=== %d. %s =============\n", list.Order, list.Name)
		a.getCardCountByWeeks(list)
	}
}

func (a *counter) getCardCountByWeeks(list models.List) {
	locKst, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Fatalf("fail to load time locale: %e", err)
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, locKst)
	oneWeekAgo := today.AddDate(0, 0, -7)
	twoWeeksAgo := today.AddDate(0, 0, -7)

	totalCount := 0
	oneWeekCount := 0
	twoWeekCount := 0
	moreAgoCount := 0
	cards := a.apiClient.GetCards(list.Id)
	for _, card := range cards {
		if !card.HasLabel(a.targetLabel) {
			continue
		}
		totalCount += 1
		act := a.apiClient.GetCreateAction(card)
		if act.Date.After(oneWeekAgo) {
			oneWeekCount += 1
		} else if act.Date.After(twoWeeksAgo) {
			twoWeekCount += 1
		} else {
			moreAgoCount += 1
		}
	}
	fmt.Printf(">>> 전체  : %d\n", totalCount)
	fmt.Printf(">>> 1주전 : %d\n", oneWeekCount)
	fmt.Printf(">>> 2주전 : %d\n", twoWeekCount)
	fmt.Printf(">>> 그이상: %d\n", moreAgoCount)
	fmt.Println(totalCount == oneWeekCount+twoWeekCount+moreAgoCount)
}
