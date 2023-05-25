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
	envStr := os.Getenv("BASE_DATE")
	if len(envStr) == 0 {
		now := time.Now()
		// 기준일자 00:00 기준이 되므로, 하루 더해줌
		now = now.AddDate(0, 0, 1)
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	}

	layout := "20060102"
	result, err := time.ParseInLocation(layout, envStr, time.Local)
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
	reqSpeed := getIntEnv("REQ_SPEED", 90)
	if reqSpeed < 0 {
		log.Fatalln("REQ_SPEED 는 0 보다 커야 합니다. (범위: 0~100)")
	}
	if reqSpeed > 100 {
		log.Fatalln("REQ_SPEED 는 100 이하여야 합니다. (trello api의 속도제한: 100요청/10초)")
	}

	baseDate := getBaseDate()
	maxWeek := getIntEnv("MAX_WEEK", -1)
	token := mustGetEnv("TOKEN")
	apiKey := mustGetEnv("API_KEY")
	boardId := mustGetEnv("BOARD_ID")
	targetLabel := mustGetEnv("LABEL")

	apiClient := api.NewTrelloClient(token, apiKey, boardId, reqSpeed)
	counter := counter.NewCounter(apiClient, targetLabel, baseDate)

	fmt.Printf("\n[1차] 현시점 '%s'라벨이 있는 카드 개수\n", targetLabel)
	printCardsCounts(counter.GetCardCounts())

	beginTime := time.Now()
	fmt.Println("\n[2차:beta] count 기간별 개수.. (3분정도 걸립니다)")
	fmt.Println("     '429 Too Many Requests' 에러가 나면, REQ_SPEED를 줄여주세요 (최대:100)")
	fmt.Printf("[BASE_DATE] %v\n", baseDate)
	weekCount := counter.GetCardCountsByWeeks()
	printCardsCountsByWeek(weekCount, maxWeek, baseDate)
	endTime := time.Now()
	fmt.Printf("\n[걸린시간] %v\n", endTime.Sub(beginTime))
}

func printCardsCounts(results models.CardCountSlice) {
	sort.Sort(results)
	for _, res := range results {
		fmt.Printf("%02d. %s: %d\n", res.Order+1, res.ListName, res.Total)
	}
}

func printCardsCountsByWeek(results models.CardCountSlice, maxWeek int, baseDate time.Time) {
	sort.Sort(results)
	for _, res := range results {
		fmt.Printf("\n=== %02d. %s: %d ===========\n", res.Order+1, res.ListName, res.Total)

		if maxWeek < 0 {
			// 정렬 없이 전부 출력
			for week, count := range res.ByWeek {
				fmt.Printf(" >>> %4d주전 생성 (~%s): %d\n", week, getWeekAgoDay(baseDate, week), count)
			}
		} else {
			beforeLimitCount := 0
			for week := 0; week < maxWeek; week++ {
				if res.ByWeek[week] > 0 {
					fmt.Printf(" >>> %3d주전 생성 (~%s): %d\n", week, getWeekAgoDay(baseDate, week), res.ByWeek[week])
					beforeLimitCount += res.ByWeek[week]
				}
			}
			fmt.Printf(" >>> %d주 보다 오래전 생성: %d\n", maxWeek, res.Total-beforeLimitCount)
		}
	}
}

func getWeekAgoDay(baseDate time.Time, week int) string {
	layout := "2006-01-02"
	ago := baseDate.AddDate(0, 0, -7*week)
	return ago.Format(layout)
}
