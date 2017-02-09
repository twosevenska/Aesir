package server

import (
	"time"

	"github.com/DeanThompson/ginpprof"
	log "github.com/Sirupsen/logrus"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// CreateRouter creates a gin Engine and adds endpoints
func CreateRouter() *gin.Engine {

	r := gin.New()
	r.Use(Logrus(log.StandardLogger()), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/forge/api", LobbyAuthentication())
	{
		//api.GET("/", controllers.INSERT_CONTROLLER_FUNCTION_HERE)
	}

	if config.Debug {
		// automatically add debugging routers
		ginpprof.Wrapper(r)
	}
	return r
}

// RunServer starts the gin Router and listens forever, recovering from panics
func RunServer() {

	// We should add the init for the forger here
	// Reminder, forger is the go file responsible for working on the forge

	r := CreateRouter()

	endless.DefaultHammerTime = 10 * time.Second
	endless.DefaultReadTimeOut = 30 * time.Second
	endless.ListenAndServe(":7001", r)
}

// Logrus midleware
func Logrus(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
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
