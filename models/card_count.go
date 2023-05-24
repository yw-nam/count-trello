package models

type CardCount struct {
	Order    int
	ListName string
	Total    int
	ByWeek   map[int]int
}

type CardCountSlice []CardCount

func (s CardCountSlice) Len() int {
	return len(s)
}

func (s CardCountSlice) Less(i, j int) bool {
	return s[i].Order < s[j].Order
}

func (s CardCountSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
