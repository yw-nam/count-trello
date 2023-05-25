package counter

import (
	"os"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yw-nam/count-trello/api"
)

var client api.Client

func TestMain(m *testing.M) {
	client = api.NewDummyClient()
	os.Exit(m.Run())
}

func TestGetCardCounts(t *testing.T) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	for _, tc := range []struct {
		isExistLabel bool
		title        string
		testCounter  *counter
	}{
		{true, "get counts", NewCounter(client, "Platform", today)},
		{false, "get counts with wrong label", NewCounter(client, "No Exist Label", today)},
	} {
		t.Run(tc.title, func(t *testing.T) {
			assert := assert.New(t)

			results := tc.testCounter.GetCardCounts()
			assert.Len(results, 2)

			sort.Sort(results)
			for i, res := range results {
				assert.Equal(i, res.Order)
			}

			if tc.isExistLabel {
				assert.Equal(2, results[0].Total)
				assert.Equal(3, results[1].Total)
			} else {
				assert.Equal(0, results[0].Total)
				assert.Equal(0, results[1].Total)
			}
		})
	}
}
