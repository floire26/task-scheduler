package middleware

import (
	"net/http"

	"github.com/floire26/task-scheduler/shared"
	"github.com/gin-gonic/gin"
)

func ErrorResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Errors.Last() != nil {

			if err, ok := c.Errors.Last().Err.(*shared.CustomError); ok {
				c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Error()})
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}

			c.Abort()
		}
	}
}
