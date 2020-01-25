package main

import (
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func parseTrainSchedules(schedule io.ReadCloser, limit int) []train {
	var trainSlice []train
	z := html.NewTokenizer(schedule)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return trainSlice

		case tt == html.StartTagToken:
			t := z.Token()

			if strings.Contains(t.String(), "table-mobile-portrait") {
				//Exit because (no idea why) they return the same results for portrait and landscape orientation
				return trainSlice
			}

			var scheduledTime string
			if strings.Contains(t.String(), "id=\"hour\"") {
				tokenType := z.Next()
				if tokenType == html.TextToken {
					scheduledTime = z.Token().Data

					//If result contains "x min", change it to HH:MM
					if strings.Contains(scheduledTime, "min") {
						scheduledTime = changeScheduleFormatToHourAndMinute(scheduledTime)
					}
				}

				for {
					if len(trainSlice) >= limit {
						break
					}

					tokenType = z.Next()
					if strings.Contains(z.Token().String(), "<span") {
						tokenType := z.Next()
						if tokenType == html.TextToken {
							destination := strings.Replace(z.Token().Data, "BARCELONA", "BCN", -1)

							//Show max 13 characters long destinations
							if len(destination) > 13 {
								destination = destination[:13]
							}

							train := newTrain(strings.TrimSpace(scheduledTime), strings.TrimSpace(destination))
							trainSlice = append(trainSlice, *train)
							break
						}
					}
				}
			}
		}
	}
}

func changeScheduleFormatToHourAndMinute(scheduledTime string) string {
	re := regexp.MustCompile("[0-9]+")
	trainMinuteString := strings.Join(re.FindAllString(scheduledTime, -1), "")
	trainMinute, _ := strconv.Atoi(trainMinuteString)
	trainSchedule := time.Now().Add(time.Duration(trainMinute) * time.Minute)

	trainScheduleHour := strconv.Itoa(trainSchedule.Hour())
	trainScheduleMinute := strconv.Itoa(trainSchedule.Minute())

	if len(trainScheduleHour) == 1 {
		trainScheduleHour = "0" + trainScheduleHour
	}

	if len(trainScheduleMinute) == 1 {
		trainScheduleMinute = "0" + trainScheduleMinute
	}

	return trainScheduleHour + ":" + trainScheduleMinute
}
