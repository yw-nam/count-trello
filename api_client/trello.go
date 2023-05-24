package api_client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/yw-nam/count-trello/models"
)

type trello struct {
	token   string
	apiKey  string
	boardId string
}

func NewTrelloApi(token, apiKey, boardId string) *trello {
	return &trello{
		token:   token,
		apiKey:  apiKey,
		boardId: boardId,
	}
}

func (a *trello) GetList() []models.List {
	lists := []models.List{}
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

func (a *trello) GetCards(listId string) []models.Card {
	cards := []models.Card{}
	urlGetCards := fmt.Sprintf("https://api.trello.com/1/lists/%s/cards?key=%s&token=%s&fields=name,labels", listId, a.apiKey, a.token)
	jsonData, err := getRespJson(urlGetCards)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(jsonData, &cards); err != nil {
		log.Fatal(err)
	}
	return cards
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
