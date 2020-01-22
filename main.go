package main

import "log"

func main() {
	schedule := getNewTrainSchedule()
	trainSchedules := parse(schedule)

	log.Println(trainSchedules)
}
