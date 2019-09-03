package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Addresss string    `form:"address""`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
	r.Run()
}

func testing(c *gin.Context) {
	var person Person

	if err := c.ShouldBind(&person); err == nil {
		c.String(http.StatusOK, "%v", person)
	} else {
		c.String(http.StatusOK, "person binding error: %v", err)
	}
}
