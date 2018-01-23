package main


type Activity struct {
	startHr int
	endHr   int
}

type Activities []Activity

func (s Activities) Len() int      { return len(s) }
func (s Activities) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByTimes struct{ Activities }

func (s ByTimes) Less(i, j int) bool {
	if s.Activities[i].startHr == s.Activities[j].startHr {
		return s.Activities[i].endHr < s.Activities[j].endHr
	}
	return s.Activities[i].startHr < s.Activities[j].startHr
}