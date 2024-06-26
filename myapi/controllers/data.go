package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var storedData [][]string

func GetData(c *gin.Context) {
	if len(storedData) == 0 {
		data := map[string]string{
			"error": "No data available",
		}
		c.JSON(http.StatusNotFound, data)
		return
	}

	c.JSON(http.StatusOK, storedData)
}
