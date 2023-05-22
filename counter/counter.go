package counter

import (
	"github.com/yw-nam/count-trello/api_client"
	"github.com/yw-nam/count-trello/model"
)

type counter struct {
	apiClient api_client.ApiClient
}

func NewCounter(api api_client.ApiClient) *counter {
	return &counter{
		apiClient: api,
	}
}

func (a *counter) GetResults() model.ResultSlice {
	lists := a.apiClient.GetList()

	ch := make(chan model.Result, len(lists))
	for i, list := range lists {
		go a.apiClient.GetResult(i, list, ch)
	}

	results := []model.Result{}
	for i := 0; i < len(lists); i++ {
		res := <-ch
		results = append(results, res)
	}
	return results
}
