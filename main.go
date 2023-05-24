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
	weekCount := counter.GetCardCountsByWeeks()
	printCardsCountsByWeek(weekCount, 8)

}

func printCardsCounts(results models.CardCountSlice) {
	sort.Sort(results)
	for _, res := range results {
		fmt.Printf("%02d. %s: %d\n", res.Order+1, res.ListName, res.Total)
	}
}

func printCardsCountsByWeek(results models.CardCountSlice, limit int) {
	sort.Sort(results)
	for _, res := range results {
		fmt.Printf("=== %02d. %s: %d ===========\n", res.Order+1, res.ListName, res.Total)

		beforeLimitCount := 0
		for week := 0; week < limit; week++ {
			if res.ByWeek[week] > 0 {
				fmt.Printf(" >>> %4d주전 생성: %d\n", week, res.ByWeek[week])
				beforeLimitCount += res.ByWeek[week]
			}
		}
		fmt.Printf(" >>> %d주 보다 오래전 생성: %d\n", limit, res.Total-beforeLimitCount)

		// // 전부 출력?
		// for week, count := range res.ByWeek {
		// 	fmt.Printf(">>> %4d주전 생성된 개수: %d\n", week, count)
		// }
	}
}
