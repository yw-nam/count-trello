package counter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type trello struct {
	token   string
	apiKey  string
	boardId string
	label   string
}

func NewCounter(token, apiKey, boardId, label string) *trello {
	return &trello{
		token:   token,
		apiKey:  apiKey,
		boardId: boardId,
		label:   label,
	}
}

func (a *trello) GetResults() ResultSlice {
	lists := a.getList()
	results := []Result{}
	ch := make(chan Result, len(lists))
	for i, list := range lists {
		go a.getResult(i, list, ch)
	}

	for i := 0; i < len(lists); i++ {
		res := <-ch
		results = append(results, res)
	}
	return results
}

func (a *trello) getResult(order int, list List, ch chan<- Result) {
	cards := []Card{}
	urlGetCards := fmt.Sprintf("https://api.trello.com/1/lists/%s/cards?key=%s&token=%s&fields=name,labels", list.Id, a.apiKey, a.token)
	jsonData, err := getRespJson(urlGetCards)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(jsonData, &cards); err != nil {
		log.Fatal(err)
	}

	count := countCardsHavingLabel(a.label, cards)
	res := Result{
		Order:     order,
		ListName:  list.Name,
		CardCount: count,
	}
	ch <- res
}

func (a *trello) getList() []List {
	lists := []List{}
	urlGetLists := fmt.Sprintf("https://api.trello.com/1/boards/%s/lists/open?key=%s&token=%s", a.boardId, a.apiKey, a.token)
	if jsonData, err := getRespJson(urlGetLists); err != nil {
		log.Fatal(err)
	} else {
		if err := json.Unmarshal(jsonData, &lists); err != nil {
			log.Fatal(err)
		}
	}
	return lists
}

func getRespJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fail to get request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response is not OK: %v", resp.Status)
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fail to read response body: %w", err)
	}
	return jsonData, nil
}
