package main

type train struct {
	schedule    string
	destination string
}

func NewTrain(sched string, dest string) *train {
	t := new(train)
	t.schedule = sched
	t.destination = dest

	return t
}
