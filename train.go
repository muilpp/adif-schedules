package main

type train struct {
	Schedule    string
	Destination string
}

func newTrain(sched string, dest string) *train {
	t := new(train)
	t.Schedule = sched
	t.Destination = dest

	return t
}

// func (t train) Schedule() string {
//  return t.Schedule
// }

// func (t train) Destination() string {
//  return t.Destination
// }
