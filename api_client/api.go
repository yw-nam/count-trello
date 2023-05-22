package api_client

import "github.com/yw-nam/count-trello/model"

type ApiClient interface {
	GetList() []model.List
	GetResult(order int, list model.List, ch chan<- model.Result)
}
