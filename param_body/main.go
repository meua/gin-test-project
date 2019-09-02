package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyBetys, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBetys))
		firstName := c.PostForm("first_name")
		lastNmae := c.DefaultPostForm("last_name", "last_default_name")

		c.String(http.StatusOK, "%s,%s,%s",
			firstName, lastNmae, string(bodyBetys))
	})
	r.Run()
}
