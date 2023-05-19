package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type list struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Closed bool   `json:"closed"`
}

type card struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Labels []label `json:"labels"`
}

func (c *card) hasLabel(targetLabel string) bool {
	for _, label := range c.Labels {
		if label.Name == targetLabel {
			return true
		}
	}
	return false
}

type label struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func mustGetEnv(key string) string {
	res := os.Getenv(key)
	if len(res) == 0 {
		log.Fatalf("%s env is required", key)
	}
	return res
}

func main() {
	token := mustGetEnv("TOKEN")
	apiKey := mustGetEnv("API_KEY")
	boardId := mustGetEnv("BOARD_ID")
	targetLabel := mustGetEnv("LABEL")

	lists := []list{}
	urlGetLists := fmt.Sprintf("https://api.trello.com/1/boards/%s/lists/open?key=%s&token=%s", boardId, apiKey, token)
	if jsonData, err := getRespJson(urlGetLists); err != nil {
		log.Fatal(err)
	} else {
		if err := json.Unmarshal(jsonData, &lists); err != nil {
			log.Fatal(err)
		}
	}

	for _, list := range lists {
		cards := []card{}
		urlGetCards := fmt.Sprintf("https://api.trello.com/1/lists/%s/cards?key=%s&token=%s&fields=name,labels", list.Id, apiKey, token)
		jsonData, err := getRespJson(urlGetCards)
		if err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(jsonData, &cards); err != nil {
			log.Fatal(err)
		}

		count := countLabel(targetLabel, cards)
		fmt.Printf("%s: %d\n", list.Name, count)
	}

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

func countLabel(label string, cards []card) int {
	count := 0
	for _, c := range cards {
		if c.hasLabel(label) {
			count += 1
		}
	}
	return count
}
