package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

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

func getBaseDate() time.Time {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatalf("fail to load time locale: %e", err)
	}

	envStr := os.Getenv("BASE_DATE")
	if len(envStr) == 0 {
		now := time.Now()
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	}

	layout := "20060102"
	result, err := time.ParseInLocation(layout, envStr, loc)
	if err != nil {
		log.Fatalf("fail to parse BASE_DATE: %v", err)
	}
	return result
}

func getIntEnv(key string, def int) int {
	envStr := os.Getenv(key)
	if len(envStr) == 0 {
		return def
	}
	result, err := strconv.Atoi(envStr)
	if err != nil {
		log.Fatalf("fail to parse int %s: %v", key, err)
	}
	return result
}

func main() {
	baseDate := getBaseDate()
	maxWeek := getIntEnv("MAX_WEEK", -1)
	token := mustGetEnv("TOKEN")
	apiKey := mustGetEnv("API_KEY")
	boardId := mustGetEnv("BOARD_ID")
	targetLabel := mustGetEnv("LABEL")

	apiClient := api.NewTrelloClient(token, apiKey, boardId)
	counter := counter.NewCounter(apiClient, targetLabel, baseDate)

	printCardsCounts(counter.GetCardCounts())

	fmt.Println("\n[2차] count 기간별 개수.. (아직 오래걸림..)")
	fmt.Printf("[BASE_DATE] %v\n", baseDate)
	weekCount := counter.GetCardCountsByWeeks()
	printCardsCountsByWeek(weekCount, maxWeek)

}

func printCardsCounts(results models.CardCountSlice) {
	sort.Sort(results)
	for _, res := range results {
		fmt.Printf("%02d. %s: %d\n", res.Order+1, res.ListName, res.Total)
	}
}

func printCardsCountsByWeek(results models.CardCountSlice, maxWeek int) {
	sort.Sort(results)
	for _, res := range results {
		fmt.Printf("=== %02d. %s: %d ===========\n", res.Order+1, res.ListName, res.Total)

		if maxWeek < 0 {
			// 정렬 없이 전부 출력
			for week, count := range res.ByWeek {
				fmt.Printf(">>> %4d주전 생성된 개수: %d\n", week, count)
			}
		} else {
			beforeLimitCount := 0
			for week := 0; week < maxWeek; week++ {
				if res.ByWeek[week] > 0 {
					fmt.Printf(" >>> %4d주전 생성: %d\n", week, res.ByWeek[week])
					beforeLimitCount += res.ByWeek[week]
				}
			}
			fmt.Printf(" >>> %d주 보다 오래전 생성: %d\n", maxWeek, res.Total-beforeLimitCount)
		}
	}
}
