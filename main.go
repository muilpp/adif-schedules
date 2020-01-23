package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// const from = "79100" //granollers centre
	// const to = "71802"   //passeig de gràcia

	// http://localhost:8080/schedules/79100/71802
	r := gin.Default()
	r.GET("/schedules/:from/:to", func(c *gin.Context) {
		from := c.Param("from")
		to := c.Param("to")

		schedule := getNewTrainSchedule(from, to)
		trainSchedules := parseTrainSchedules(schedule)

		// for _, schedule := range trainSchedules {
		//  log.Printf("Train destination to %v leaves at %v", schedule.Destination, schedule.Schedule)
		// }
		c.JSON(200, trainSchedules)
	})

	//Allow all origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.Run()
}
