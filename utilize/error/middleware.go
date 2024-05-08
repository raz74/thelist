package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		switch err.Err {
		case RestaurantNotFoundError:
			c.JSON(http.StatusNotFound, gin.H{"message": RestaurantNotFoundError.Error()})
		case RestaurantDontCreatedError:
			c.JSON(http.StatusInternalServerError, gin.H{"message": RestaurantDontCreatedError.Error()})
		case SendNotificationFailedError:
			c.JSON(http.StatusInternalServerError, gin.H{"message": SendNotificationFailedError.Error()})
		default:
			c.JSON(http.StatusInternalServerError, "encountered internal server error")
		}
	}
}
