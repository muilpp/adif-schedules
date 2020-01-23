package main

import (
	"io"
	"log"
	"net/http"
)

func getNewTrainSchedule(from string, to string) io.ReadCloser {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://elcanoweb.adif.es/departures/list?station="+from+"+&dest="+to+"&previous=1&showCercanias=true&showOtros=false&isNative=false", nil)
	req.SetBasicAuth("deimos", "deimostt")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return resp.Body
}
