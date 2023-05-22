package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/yw-nam/count-trello/api_client"
	"github.com/yw-nam/count-trello/counter"
	"github.com/yw-nam/count-trello/model"
)

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

	trelloApi := api_client.NewTrelloApi(token, apiKey, boardId, targetLabel)
	counter := counter.NewCounter(trelloApi)

	result := counter.GetResults()
	printCardsCount(result)
}

func printCardsCount(results model.ResultSlice) {
	sort.Sort(results)
	for _, res := range results {
		fmt.Printf("%02d. %s: %d\n", res.Order+1, res.ListName, res.CardCount)
	}
}
