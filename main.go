package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/yw-nam/count-trello/api"
	"github.com/yw-nam/count-trello/counter"
	"github.com/yw-nam/count-trello/models"
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

	apiClient := api.NewTrelloClient(token, apiKey, boardId)
	counter := counter.NewCounter(apiClient, targetLabel)

	printCardsCounts(counter.GetCardCounts())

	fmt.Println("\n[2차] count 기간별 개수.. (아직 오래걸림..)")
	counter.GetCardCountsByWeeks()
}

func printCardsCounts(results models.CardCountSlice) {
	sort.Sort(results)
	for _, res := range results {
		fmt.Printf("%02d. %s: %d\n", res.Order+1, res.ListName, res.CardCount)
	}
}
