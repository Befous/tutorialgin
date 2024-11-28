package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ContohHandler(c *gin.Context) {
	zParam := c.Param("z")
	var x int = 2
	var y int = 2
	var sum int = x + y
	z, err := strconv.Atoi(zParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "z harus diisi dengan angka",
		})
		return
	}
	if z != sum {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Jawaban salah",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Jawaban benar",
	})
}
