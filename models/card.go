package models

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

func CountCardsHavingLabel(label string, cards []Card) int {
	count := 0
	for _, c := range cards {
		if c.hasLabel(label) {
			count += 1
		}
	}
	return count
}
