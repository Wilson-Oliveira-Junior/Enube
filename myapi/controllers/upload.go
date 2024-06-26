package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	tempFile := "./" + file.Filename
	if err := c.SaveUploadedFile(file, tempFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	xlsFile, err := xlsx.OpenFile(tempFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read file"})
		return
	}

	var newData [][]string
	for _, sheet := range xlsFile.Sheets {
		for _, row := range sheet.Rows {
			var values []string
			for _, cell := range row.Cells {
				values = append(values, cell.String())
			}
			newData = append(newData, values)
		}
	}

	storedData = newData

	os.Remove(tempFile)

	c.JSON(http.StatusOK, gin.H{"message": "File processed successfully"})
}
