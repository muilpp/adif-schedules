package main

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func parse(schedule io.ReadCloser) []train {
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
				}

				for {
					tokenType = z.Next()
					if strings.Contains(z.Token().String(), "<span") {
						tokenType := z.Next()
						if tokenType == html.TextToken {
							train := NewTrain(scheduledTime, z.Token().Data)
							//log.Println("Found: ", *train)
							trainSlice = append(trainSlice, *train)
							break
						}
					}
				}
			}
		}
	}
}
