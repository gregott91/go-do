package godo

type byAge []*Note

func (s byAge) Len() int {
	return len(s)
}
func (s byAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byAge) Less(i, j int) bool {
	return s[i].Timestamp.After(s[j].Timestamp)
}
