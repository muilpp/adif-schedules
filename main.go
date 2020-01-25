package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// "79100" -> granollers centre
// "71802" -> passeig de gràcia
// http://localhost:8080/schedules/79100/71802

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/schedules/:from/:to", func(c *gin.Context) {
		from := c.Param("from")
		to := c.Param("to")
		limit, _ := strconv.Atoi(c.Query("limit"))

		schedule := getNewTrainSchedule(from, to)
		trainSchedules := parseTrainSchedules(schedule, limit)
		c.JSON(http.StatusOK, trainSchedules)
	})

	r.Run(":8181")
}
