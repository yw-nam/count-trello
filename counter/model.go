package counter

type Result struct {
	Order     int
	ListName  string
	CardCount int
}

type ResultSlice []Result

func (s ResultSlice) Len() int {
	return len(s)
}

func (s ResultSlice) Less(i, j int) bool {
	return s[i].Order < s[j].Order
}

func (s ResultSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Label struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type List struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Closed bool   `json:"closed"`
}

type Card struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
}

func (c *Card) hasLabel(targetLabel string) bool {
	for _, label := range c.Labels {
		if label.Name == targetLabel {
			return true
		}
	}
	return false
}

func countCardsHavingLabel(label string, cards []Card) int {
	count := 0
	for _, c := range cards {
		if c.hasLabel(label) {
			count += 1
		}
	}
	return count
}
