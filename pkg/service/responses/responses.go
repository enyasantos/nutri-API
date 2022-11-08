package responses

import (
	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

func Error(c *gin.Context, statusCode int, error error) {
	JSON(c, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: error.Error(),
	})
}
