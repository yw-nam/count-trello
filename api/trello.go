package api

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

func NewTrelloClient(token, apiKey, boardId string) *trello {
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
	urlGetCards := fmt.Sprintf("https://api.trello.com/1/lists/%s/cards?key=%s&token=%s&fields=id,name,labels", listId, a.apiKey, a.token)
	jsonData, err := getRespJson(urlGetCards)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(jsonData, &cards); err != nil {
		log.Fatal(err)
	}
	return cards
}

func (a *trello) GetCreateAction(card models.Card) models.Action {
	createActs := a.getActions(card.Id, models.ActionCreateCard)
	if len(createActs) == 1 {
		return createActs[0]
	}
	if len(createActs) > 1 {
		log.Fatalf("creation action happen more than once: %d (card_id:%v)", len(createActs), card.Id)
	}

	// 없으면 가장 오래된 액션을 찾기..
	lastAct, err := a.getLastActions(card.Id)
	if err != nil {
		log.Fatal(err)
	}
	return lastAct
}

func (a *trello) getLastActions(cardId string) (models.Action, error) {
	sizeOfPage := 50
	result := models.Action{}
	urlBase := fmt.Sprintf("https://api.trello.com/1/cards/%s/actions?key=%s&token=%s&filter=%s", cardId, a.apiKey, a.token, "all")
	limit := 100 // 100페이지 넘는건 에러일거 같아서 종료
	for page := 0; page < limit; page++ {
		results := []models.Action{}
		url := fmt.Sprintf("%s&page=%d", urlBase, page)
		jsonData, err := getRespJson(url)
		if err != nil {
			return result, err
		}
		if err := json.Unmarshal(jsonData, &results); err != nil {
			return result, err
		}

		if len(results) == 0 {
			if len(result.Id) == 0 {
				return result, fmt.Errorf("cannot find action of card: %s", cardId)
			}
			return result, nil
		}
		if len(results) < sizeOfPage {
			return results[len(results)-1], nil
		}
		if len(results) >= sizeOfPage {
			result = results[len(results)-1]
		}
	}
	return result, fmt.Errorf("too big page: maybe something wrong: %s", cardId)
}

func (a *trello) getActions(cardId string, actionType string) []models.Action {
	results := []models.Action{}
	url := fmt.Sprintf("https://api.trello.com/1/cards/%s/actions?key=%s&token=%s&filter=%s", cardId, a.apiKey, a.token, actionType)
	jsonData, err := getRespJson(url)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(jsonData, &results); err != nil {
		log.Fatal(err)
	}
	return results
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
