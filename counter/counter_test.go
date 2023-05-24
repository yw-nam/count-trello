package counter

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yw-nam/count-trello/api_client"
)

var dummyApi api_client.ApiClient

func TestMain(m *testing.M) {
	dummyApi = api_client.NewDummyApi()
	os.Exit(m.Run())
}

func TestGetCardCounts(t *testing.T) {
	for _, tc := range []struct {
		isExistLabel bool
		title        string
		testCounter  *counter
	}{
		{true, "get counts", NewCounter(dummyApi, "Platform")},
		{false, "get counts with wrong label", NewCounter(dummyApi, "No Exist Label")},
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
				assert.Equal(2, results[0].CardCount)
				assert.Equal(3, results[1].CardCount)
			} else {
				assert.Equal(0, results[0].CardCount)
				assert.Equal(0, results[1].CardCount)
			}
		})
	}
}
