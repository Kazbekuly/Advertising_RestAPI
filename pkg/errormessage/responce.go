package errormessage

import (
	"Advertising/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errors struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errors{Message: message})
	logger := logging.GetLogger()
	logger.Info("error handler")
}
