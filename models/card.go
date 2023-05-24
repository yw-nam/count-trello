package models

type Card struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
}

func (c *Card) HasLabel(targetLabel string) bool {
	for _, label := range c.Labels {
		if label.Name == targetLabel {
			return true
		}
	}
	return false
}
