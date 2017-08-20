package server

import (
	"time"

	"github.com/DeanThompson/ginpprof"
	log "github.com/Sirupsen/logrus"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/twosevenska/aesir/forge/mongo"
)

// Config is a populated by env variables and Vault
type Config struct {
	Debug         bool     `envconfig:"debug" default:"false"`
	MongoHosts    []string `envconfig:"mongo_hosts" default:"mongo.dev:27017"`
	MongoDBName   string   `envconfig:"mongo_dbname" default:"forge"`
	MongoUser     string   `envconfig:"mongo_user" default:"forge"`
	MongoPassword string   `envconfig:"mongo_password" default:"forge1234"`
}

// ContextParams holds the objects required to initialise the server
type ContextParams struct {
	Config      Config
	MongoClient *mongo.Client
}

// Run starts the gin Router and listens forever, recovering from panics
func Run(c Config, mc *mongo.Client) {

	contextParams := ContextParams{
		Config:      c,
		MongoClient: mc,
	}

	r := CreateRouter(&contextParams)

	endless.DefaultHammerTime = 10 * time.Second
	endless.DefaultReadTimeOut = 30 * time.Second
	endless.ListenAndServe(":7001", r)
}

// CreateRouter creates a gin Engine and adds endpoints
func CreateRouter(contextParams *ContextParams) *gin.Engine {

	//TODO: Add custom validator later on
	r := gin.New()
	r.Use(Logrus(log.StandardLogger()), gin.Recovery())

	r.Use(ContextObjects(contextParams))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/*api := r.Group("/forge/api", Auth())
	{
		//api.GET("/", controllers.INSERT_CONTROLLER_FUNCTION_HERE)
	}*/

	if contextParams.Config.Debug {
		// automatically add routers for net/http/pprof
		// e.g. /debug/pprof, /debug/pprof/heap, etc.
		ginpprof.Wrapper(r)
	}
	return r
}

// ContextObjects attaches backend clients to the API context
func ContextObjects(contextParams *ContextParams) gin.HandlerFunc {

	return func(c *gin.Context) {
		newMongo := contextParams.MongoClient.Copy()
		defer newMongo.Close()

		c.Set("mongo", newMongo)
		c.Next()
	}
}

// Logrus midleware
func Logrus(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		c.Next()

		entry := logger.WithFields(log.Fields{
		//add fields
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		} else {
			entry.Infof("%s %s", c.Request.Method, path)
		}
	}
}
