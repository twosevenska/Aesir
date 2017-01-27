package server

import (
	"aesir/mongo"
	"time"
)

func RunServer(conf Config, mongo mongo.Client) {
	r := CreateRouter(mongo, producer, nflex, babel)
	r.Run(":8012")
}

// Logrus midleware
func Logrus(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		entry := logger.WithFields(log.Fields{
			"request": &request,
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		} else {
			entry.Infof("%s %s", c.Request.Method, path)
		}
	}
}
