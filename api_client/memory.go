package api_client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yw-nam/count-trello/models"
)

const (
	listId01 string = "559fb0c36ef933f0ca0e2729"
	listId02 string = "5f433946d90bd37728a66e61"
)

type dummy struct {
	respGetList        string
	respGetCardsByList map[string]string
}

func NewDummyApi() *dummy {
	return &dummy{
		respGetList: fmt.Sprintf(respListFmt, listId01, listId02),
		respGetCardsByList: map[string]string{
			listId01: respList01,
			listId02: respList02,
		},
	}
}

func (a *dummy) GetList() []models.List {
	lists := []models.List{}
	jsonData := []byte(a.respGetList)
	if err := json.Unmarshal(jsonData, &lists); err != nil {
		log.Fatal(err)
	}
	return lists
}

func (a *dummy) GetCards(listId string) []models.Card {
	cards := []models.Card{}
	if jsonData, isExist := a.respGetCardsByList[listId]; isExist {
		if err := json.Unmarshal([]byte(jsonData), &cards); err != nil {
			log.Fatal(err)
		}
	}
	return cards
}

const respListFmt string = `[
    {
        "id": "%s",
        "name": "고객 요구사항",
        "closed": false,
        "idBoard": "559f6a6380805f9f6597668c",
        "pos": 32767,
        "subscribed": false,
        "softLimit": null,
        "status": null
    },
    {
        "id": "%s",
        "name": "고객 확인 필요",
        "closed": false,
        "idBoard": "559f6a6380805f9f6597668c",
        "pos": 106495,
        "subscribed": false,
        "softLimit": null,
        "status": null
    }]`

const (
	respList01 string = `[
		{
			"id": "5722e11784a354345525e0c3",
			"name": "고객 요구사항 작성시 참고사항",
			"labels": [
				{
					"id": "63048950cd8865008d1a29a2",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "공지사항",
					"color": null
				}
			]
		},
		{
			"id": "642b921b2cda039bbc3b7519",
			"name": "이슈1",
			"labels": []
		},
		{
			"id": "6451b6ff63ffbacd1ec37c4b",
			"name": "{DEV} 이슈2",
			"labels": []
		},
		{
			"id": "646302889679dc6a24bbc433",
			"name": "{DEV} 이슈3",
			"labels": [
				{
					"id": "559f6a6319ad3a5dc2c3b21f",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "Platform",
					"color": "orange"
				}
			]
		},
		{
			"id": "64192d99973c526facd0ebef",
			"name": "이슈4",
			"labels": [
				{
					"id": "5f6c36807f3994273966dcfe",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "KSH",
					"color": "green"
				},
				{
					"id": "55ac4cad19ad3a5dc2d90f0e",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "웹테크",
					"color": "sky"
				},
				{
					"id": "630483ce87770200878f8663",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "VOD",
					"color": "red"
				}
			]
		},
		{
			"id": "646c26c976156b957e5989f4",
			"name": "이슈5",
			"labels": [
				{
					"id": "559f6a6319ad3a5dc2c3b21f",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "Platform",
					"color": "orange"
				},
				{
					"id": "63a3ec09b1b6a804c393e649",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "YCH",
					"color": "green"
				}
			]
		},
		{
			"id": "646c676c0281944fc3b10ac3",
			"name": "이슈6",
			"labels": []
		}
	]`

	respList02 string = `[
		{
			"id": "6439267b7ba356c36e33b5af",
			"name": "dummy 요청",
			"labels": []
		},
		{
			"id": "643f333e74d3e4a86c890209",
			"name": "dummy 전달",
			"labels": [
				{
					"id": "559f6a6319ad3a5dc2c3b221",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "Player",
					"color": "purple"
				},
				{
					"id": "5f7bc6679aeb0d34bc3e4818",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "KYK",
					"color": "green"
				}
			]
		},
		{
			"id": "640fc4c27c1edbfebf73a741",
			"name": "dummy 기능",
			"labels": []
		},
		{
			"id": "64366e8c7f6e8e507d6e5a5c",
			"name": "dummy API",
			"labels": []
		},
		{
			"id": "64366fe934c3dbcc692dc22b",
			"name": "dummy API",
			"labels": []
		},
		{
			"id": "64225833fe163c60d113a644",
			"name": "dummy 요청",
			"labels": []
		},
		{
			"id": "640ff8c1277cedcd15ba8297",
			"name": "dummy 통일",
			"labels": [
				{
					"id": "559f6a6319ad3a5dc2c3b221",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "Player",
					"color": "purple"
				}
			]
		},
		{
			"id": "64056743e7acaa87fd6e8e70",
			"name": "dummy 요청",
			"labels": [
				{
					"id": "55ac4cad19ad3a5dc2d90f0e",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "웹테크",
					"color": "sky"
				},
				{
					"id": "630483ce87770200878f8663",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "VOD",
					"color": "red"
				},
				{
					"id": "5f6c36807f3994273966dcfe",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "KSH",
					"color": "green"
				}
			]
		},
		{
			"id": "634df702f2bd5800b52ceccf",
			"name": "{DEV} dummy",
			"labels": [
				{
					"id": "55ac4cad19ad3a5dc2d90f0e",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "웹테크",
					"color": "sky"
				},
				{
					"id": "630483ce87770200878f8663",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "VOD",
					"color": "red"
				}
			]
		},
		{
			"id": "63509751fa375f0311e839c3",
			"name": "dummy 요청",
			"labels": [
				{
					"id": "55ac4cad19ad3a5dc2d90f0e",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "웹테크",
					"color": "sky"
				},
				{
					"id": "630483dc77186c00d9262e89",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "LIVE",
					"color": "blue"
				}
			]
		},
		{
			"id": "63d88c8b262361eae5d15d8f",
			"name": "dummy 요청",
			"labels": []
		},
		{
			"id": "63ec2f258a867bcd824be58d",
			"name": "dummy 요청",
			"labels": [
				{
					"id": "5f7bc6679aeb0d34bc3e4818",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "KYK",
					"color": "green"
				},
				{
					"id": "559f6a6319ad3a5dc2c3b221",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "Player",
					"color": "purple"
				}
			]
		},
		{
			"id": "5d3505a400fe345784d33caf",
			"name": "dummy 요청",
			"labels": [
				{
					"id": "5f6c36807f3994273966dcfe",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "KSH",
					"color": "green"
				},
				{
					"id": "55ac4cad19ad3a5dc2d90f0e",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "웹테크",
					"color": "sky"
				},
				{
					"id": "630483dc77186c00d9262e89",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "LIVE",
					"color": "blue"
				},
				{
					"id": "5ece0ed7bf01ca06b23534e6",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "데이터플랫폼",
					"color": "yellow"
				}
			]
		},
		{
			"id": "5b46885ee7df937d22c734cc",
			"name": "dummy 지원",
			"labels": [
				{
					"id": "559f6a6319ad3a5dc2c3b221",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "Player",
					"color": "purple"
				},
				{
					"id": "559f6a6319ad3a5dc2c3b21f",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "Platform",
					"color": "orange"
				},
				{
					"id": "55ac4cad19ad3a5dc2d90f0e",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "웹테크",
					"color": "sky"
				},
				{
					"id": "5f6c36807f3994273966dcfe",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "KSH",
					"color": "green"
				}
			]
		},
		{
			"id": "6334e6aac78d7c00a1bb60d0",
			"name": "{DEV} dummy 필요",
			"labels": []
		},
		{
			"id": "62ea0ad955fe1c29d34d68aa",
			"name": "dummy 건",
			"labels": []
		},
		{
			"id": "60235022f1935f19224f943c",
			"name": "dummy 현상",
			"labels": [
				{
					"id": "55ac4cad19ad3a5dc2d90f0e",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "웹테크",
					"color": "sky"
				}
			]
		},
		{
			"id": "61cd50fdfce991519370b0f6",
			"name": "요청",
			"labels": []
		},
		{
			"id": "60f765b5c01df011acff99ae",
			"name": "{DEV} API",
			"labels": [
				{
					"id": "5f6c3576fe57a4507ded3533",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "강진욱",
					"color": "green"
				},
				{
					"id": "559f6a6319ad3a5dc2c3b21f",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "Platform",
					"color": "orange"
				},
				{
					"id": "5ece0ed7bf01ca06b23534e6",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "데이터플랫폼",
					"color": "yellow"
				}
			]
		},
		{
			"id": "642fc73299d6a146d1405435",
			"name": "서비스 점검",
			"labels": [
				{
					"id": "559f6a6319ad3a5dc2c3b21f",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "Platform",
					"color": "orange"
				},
				{
					"id": "612ee51db115cb214a89b91d",
					"idBoard": "559f6a6380805f9f6597668c",
					"name": "PDY",
					"color": "green"
				}
			]
		}
	]`
)
