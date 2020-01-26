package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/schedules/:from/:to", func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.Query("limit"))

		schedule := getNewTrainSchedule(c.Param("from"), c.Param("to"))
		trainSchedules := parseTrainSchedules(schedule, limit)
		c.JSON(http.StatusOK, trainSchedules)
	})

	r.Run(":8181")
}
